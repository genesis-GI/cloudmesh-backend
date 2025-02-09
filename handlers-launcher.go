package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getVersions(c *gin.Context){
	email := c.Param("email")

	var account Account
	err := db.Collection("accounts").FindOne(context.TODO(), bson.M{"email": email}).Decode(&account)
	if err == mongo.ErrNoDocuments {
		c.JSON(404, gin.H{"error": "Account not found"})
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	var versionsCollection *mongo.Collection = client.Database("genesis").Collection("versions")
	var result bson.M
	err = versionsCollection.FindOne(context.TODO(), bson.M{}).Decode(&result)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch versions"})
		return
	}

	builds := result["builds"].(primitive.A)
	allowedBuilds := []interface{}{}

	for _, build := range builds {
		buildMap := build.(primitive.M)
		requiredWaveAccess := int(buildMap["requiredWaveAccess"].(int32))
		if account.Wave <= requiredWaveAccess {
			allowedBuilds = append(allowedBuilds, buildMap)
		}
	}

	c.JSON(200, gin.H{"builds": allowedBuilds})
}