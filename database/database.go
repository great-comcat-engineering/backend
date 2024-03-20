package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"greatcomcatengineering.com/backend/configs"
	"log"
	"time"
)

const DATABASE_NAME = "gce-backend"
const COLLECTION_USERS = "users"
const COLLECTION_PRODUCTS = "products"
const COLLECTION_ORDERS = "orders"

var Client *mongo.Client
var isConnected bool = false

func ConnectToMongoDB() {
	// Check if we already have an established connection
	if Client != nil && isConnected {
		log.Println("Already connected to MongoDB.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.AppConfig().Database.MongoUri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	// Ping the database to verify connection is successful
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB: ", err)
	}

	log.Println("Connected to MongoDB!")
	Client = client
	isConnected = true // Mark the connection as established
}

func DisconnectFromMongoDB() {
	if Client != nil && isConnected {
		err := Client.Disconnect(context.Background())
		if err != nil {
			log.Fatal("Failed to disconnect from MongoDB: ", err)
		}
		log.Println("Disconnected from MongoDB.")
		isConnected = false // Reset the connection flag
	}
}
