package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL not set in .env file")
	}

	const dbName = "netflix"
	const colName = "watchlist"

	clientOptions := options.Client().ApplyURI(dbUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}
	fmt.Println("MongoDB connection established successfully")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection reference is ready")
}

func DBConnection() *mongo.Collection {
	return collection
}
