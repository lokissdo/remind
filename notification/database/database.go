package database

import (
	"context"
	"log"
	"remind/notification/common/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var FCMTokenCollection *mongo.Collection

func Init() {
	InitMongoDB()
}

func InitMongoDB() {
	var err error

	// Set client options
	clientOptions := options.Client().ApplyURI(configs.Yaml.MongoDB.Connection) // Replace with your MongoDB URI

	// Connect to MongoDB
	MongoClient, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = MongoClient.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	err = MongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	// Set up the collection
	FCMTokenCollection = MongoClient.Database("notification").Collection("fcmtokens")
	log.Println("Connected to MongoDB and set up collection")
}
