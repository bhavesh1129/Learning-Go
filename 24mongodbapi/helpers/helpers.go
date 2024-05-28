package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/bhavesh1129/mongodbapi/database"
	"github.com/bhavesh1129/mongodbapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// insert one record
func InsertOneMovie(movie models.Netflix) {
	collections := database.DBConnection()
	data, err := collections.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in the db with the id: ", data.InsertedID)
}

// update one record
func UpdateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	collections := database.DBConnection()
	result, err := collections.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}

// delete one record
func DeleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	collections := database.DBConnection()
	deleteCount, err := collections.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count: ", deleteCount)
}

// delete all record
func DeleteAllMovies() {
	collections := database.DBConnection()
	deleteCounts, err := collections.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count: ", deleteCounts.DeletedCount)
}

// get all movies
func GetAllMovies() []primitive.M {
	collections := database.DBConnection()
	cur, err := collections.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}
