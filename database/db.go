package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"sync"
)

const DATABASE_NAME = "gce-backend"
const COLLECTION_USERS = "users"
const COLLECTION_PRODUCTS = "products"
const COLLECTION_ORDERS = "orders"

var (
	client     *mongo.Client
	clientOnce sync.Once
)

// GetClient provides a singleton MongoDB client
func GetClient() (*mongo.Client, error) {
	var err error
	clientOnce.Do(func() {
		mongoURI := os.Getenv("MONGODB_URI")
		if mongoURI == "" {
			err = fmt.Errorf("MONGODB_URI environment variable not set")
			return
		}

		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPIOptions)

		client, err = mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			err = fmt.Errorf("failed to connect to MongoDB: %w", err)
			return
		}

		if err = client.Ping(context.Background(), nil); err != nil {
			err = fmt.Errorf("failed to ping MongoDB: %w", err)
			return
		}

		fmt.Println("Successfully connected to MongoDB")
	})

	return client, err
}
