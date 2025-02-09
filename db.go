package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func isValidUsername(username string) bool {
    // Define a regex pattern for valid usernames (e.g., alphanumeric characters)
    var validUsernamePattern = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
    return validUsernamePattern.MatchString(username)
}

var db *mongo.Database
var client *mongo.Client

func initDB() error {

	var clientOptions *options.ClientOptions

	

	if !useRemoteDB {
		color.Cyan("[INFO]: Connecting to local db...")
		clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	}else{
		color.Cyan("[INFO]: Connecting to remote db...")
		clientOptions = options.Client().ApplyURI("mongodb://81.10.229.31:38128")
	}

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		color.Red("[✗ FAILURE] Failed to connect to the database")
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		color.Red("[✗ FAILURE] Database is unavailable")
		return err
	}

	rwEnv := os.Getenv("RAILWAY_ENVIRONMENT")
	if rwEnv == "production" {
		db = client.Database("genesis")
		color.Cyan("[INFO]: Connecting to production db...")
	}else{
		db = client.Database("genesis-development")
		color.Cyan("[INFO]: Connecting to development db...")
	}

	color.Green("[✓ SUCCESS] Connected to DB successfully!")
	return nil
}



func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func findByEmail(email string) (bool, string) {
	if !isValidEmail(email) {
		log.Println("Invalid email format")
		return false, ""
	}
	filter := bson.M{"email": email}

	var result bson.M
	err := db.Collection("accounts").FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, ""
	} else if err != nil {
		//log.Printf("Error occurred while finding email: %v\n", err)
		return false, ""
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error converting document to JSON: %v\n", err)
		return false, ""
	}

	return true, string(jsonData)
}

func login(req LoginRequest) (bool, string) {
	// Check if email is in a valid format -> No 'SQL' injection
	if !isValidEmail(req.Email) {
		return false, "Invalid email or password"
	}
	req.Email = strings.TrimSpace(req.Email)


	var account Account
	filter := bson.D{{Key: "email", Value: req.Email}}
	err := db.Collection("accounts").FindOne(context.TODO(), filter).Decode(&account)

	// Password checks
	if err == mongo.ErrNoDocuments {

		return false, "Invalid email or password"
	} else if err != nil {

		return false, "An error occurred. Please try again later."
	}

	// Check if passwords match
	if !CheckPasswordHash(req.Password, account.Password) {
		return false, "Invalid email or password"
	}

	color.Cyan("[INFO]: Login successful for email: %s", req.Email)
	// Login successful
	return true, "Login successful!"
}



func register(email, username, password string) (bool, string) {
    // Check email format
    if !isValidEmail(email) {
        return false, "Invalid email format"
    }

    // Check if email is already in use
    foundEmail, _ := findByEmail(email)
    if foundEmail {
        return false, "Email already in use"
    }

    // Check if username is already in use
    foundUsername, _ := findByUsername(username)
    if foundUsername {
        return false, "Username already in use"
    }

    hashedPW, err := HashPassword(password)
    if err != nil {
        return false, "Error: " + err.Error()
    }

    newAccount := Account{
        Email:     email,
        Username:  username,
        Password:  hashedPW,
        Admin:     false,
        Wave:      5,
        CreatedAt: time.Now(),
    }

    _, creationErr := db.Collection("accounts").InsertOne(context.TODO(), newAccount)
    if creationErr != nil {
        return false, "Failed to register account!"
    }
    return true, "Account registered successfully!"
}

func findByUsername(username string) (bool, string) {
    if !isValidUsername(username) {
        return false, "Invalid username format"
    }
    filter := bson.M{"username": username}
    var result bson.M
    err := db.Collection("accounts").FindOne(context.TODO(), filter).Decode(&result)
    if err == mongo.ErrNoDocuments {
        return false, ""
    } else if err != nil {
        return false, ""
    }
    jsonData, err := json.Marshal(result)
    if err != nil {
        log.Printf("Error converting document to JSON: %v\n", err)
        return false, ""
    }
    return true, string(jsonData)
}

func setMOTD(text string, currentTime time.Time) (bool, string) {
	

	filter := bson.M{"_id": "motd"}
	update := bson.M{
		"$set": bson.M{
			"text": text,
			"time": currentTime,
		},
	}


	opts := options.Update().SetUpsert(true)
	_, err := db.Collection("motd").UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return false, "Failed to update MOTD: " + err.Error()
	}
	return true, "MOTD updated successfully!"
}


func getMOTD() (string, string) {
	filter := bson.M{"_id": "motd"}

	var result struct {
		Text string    `bson:"text"`
		Time time.Time `bson:"time"`
	}

	err := db.Collection("motd").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Printf("Failed to retrieve MOTD: %v", err)
		return "", ""
	}

	// Format the time as a string, for example in RFC3339 format.
	return result.Text, result.Time.Format(time.RFC3339)
}
