package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}

// MongoConfig is the configuration for the MongoDB database

type MongoConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

// MongoConnection is the connection to the MongoDB database

type MongoConnection struct {
	Client *mongo.Client
}

// NewMongoConnection creates a new connection to the MongoDB database

func NewMongoConnection(config MongoConfig) (*MongoConnection, error) {
	// Create a new connection to the MongoDB database
	// ...
	fmt.Print("MongoDB connection established")
	return &MongoConnection{}, nil
}

// Close closes the connection to the MongoDB database

func (c *MongoConnection) Close() error {
	// Close the connection to the MongoDB database
	// ...
	fmt.Print("MongoDB connection closed")

	return nil
}
