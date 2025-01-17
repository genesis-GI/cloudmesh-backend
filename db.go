package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"

)

import (
	"regexp"
)

var accounts *mongo.Collection

func initDB(){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	accounts = client.Database("genesis").Collection("accounts")
	fmt.Println("DB initiated successfully")
}


func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func findByEmail(email string) (bool, string) {
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


func isValidEmail(email string) bool {
	// Regular expression for validating email format
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func login(req LoginRequest) (bool, string){
	if !isValidEmail(req.Email) {
		return false, "Invalid email format"
	}
	var account Account
	err := accounts.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&account)


	if err == mongo.ErrNoDocuments{
		return false, "Account not found! Please register first"
	} else if err != nil{
		return false, "Error: " + err.Error()
	}


	if !CheckPasswordHash(req.Password, account.Password) {
		return false, "Invalid email or password"
	}
	return true, "Login successful!"
}


func register(email, username, password string) (bool, string){
	found, _ := findByEmail(email)
	if found{
		return false, "User already registered!"
	}

	hashedPW, err := HashPassword(password)
	// fmt.Println("password entered:",password)
	// fmt.Println("Hashed pw:", hashedPW)
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