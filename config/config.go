package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoURI    string
	JWTSecret   string
	MongoClient *mongo.Client
)

func LoadConfig() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve environment variables
	MongoURI = os.Getenv("MONGODB_URI")
	JWTSecret = os.Getenv("JWT_SECRET")

	if MongoURI == "" || JWTSecret == "" {
		log.Fatal("Environment variables MONGODB_URI or JWT_SECRET not set")
	}

	// Initialize MongoDB client with context and connection options
	clientOptions := options.Client().ApplyURI(MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	MongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Verify the connection
	err = MongoClient.Ping(ctx, nil)

	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	} else {
		fmt.Println("Connected to MongoDB successfully")
	}
}
