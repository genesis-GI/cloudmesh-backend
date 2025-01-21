package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"regexp"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var accounts *mongo.Collection

func initDB() error {

	var clientOptions *options.ClientOptions
	if gin.ReleaseMode == gin.DebugMode {
		clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	}else{
		clientOptions = options.Client().ApplyURI("mongodb://81.10.229.31:27017")
	}
	

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("database is unavailable: %w", err)
	}

	accounts = client.Database("genesis").Collection("accounts")
	fmt.Println("DB initiated successfully")
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
	err := accounts.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, ""
	} else if err != nil {
		log.Printf("Error occurred while finding email: %v\n", err)
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
	err := accounts.FindOne(context.TODO(), filter).Decode(&account)

	// Password checks
	if err == mongo.ErrNoDocuments {
		log.Printf("Login attempt for non-existent email: %s\n", req.Email)
		return false, "Invalid email or password"
	} else if err != nil {
		log.Printf("Database error during login for email %s: %v\n", req.Email, err)
		return false, "An error occurred. Please try again later."
	}

	// Check if passwords match
	if !CheckPasswordHash(req.Password, account.Password) {

		log.Printf("Invalid password attempt for email: %s\n", req.Email)
		return false, "Invalid email or password"
	}

	// Login successful
	return true, "Login successful!"
}



func register(email, username, password string) (bool, string){
	found, _ := findByEmail(email)
	if found{
		return false, "User already registered!"
	}

	hashedPW, err := HashPassword(password)
	if err != nil{
		return false, "Error: "+err.Error()
	}

	newAccount := Account{
		Email: email,
		Username: username,
		Password: hashedPW,
		Admin: false,
		Wave: 5,
		CreatedAt: time.Now(),
	}
	
	_, creationErr := accounts.InsertOne(context.TODO(), newAccount)

	if creationErr != nil{
		return false, "Failed to register account!"
	}

	return true, "Account registered successfully!"
}