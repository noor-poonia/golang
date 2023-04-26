// database connection, helper and controller function all in this file

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noor-poonia/mongoApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb+srv://golang:noor@cluster0.railjru.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"    //database name
const colName = "watchlist" //collection name

// most important
var collection *mongo.Collection

//connect with MongoDB
// init() function is the first function to be executed when the program runs and its executed only once
// it is basically used for initialisation 
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	// Package context defines the Context type, which carries deadlines, cancellation signals, and
	// other request-scoped values across API boundaries and between processes.
	// confusing topic -> TODO is type of context
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection successful.")

	//collection instance
	// collection = (*mongo.Collection)(client.Database(dbName))
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("collection instance is ready")
}

// MongoDB helper -file

//insert 1 record -> helper method -> not exported
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		// more control version of panic
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in db with id:", inserted.InsertedID)
}

// update 1 movie 
func updateOneMovie(movieId string)  {
	// ObjectIDFromHex -> converts string into acceptible mongodb id (_id) format 
	id, _ := primitive.ObjectIDFromHex(movieId)
	// data in mongodb is in bson format 
	// bson.M{} -> short and clear, lowercase capital no difference -> majority use case -> unordered representation 
	// bson.D{} -> ordered representaiton
	// provide query parameters
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", result.ModifiedCount)
}

// delete 1 movie 
func deleteOneMovie(movieId string)  {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete count:", deleteCount)
}

// delete all records from mongodb 
func deleteAllMovies() int64 {
	// no value provided -> everything is included
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete count:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from db
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// actual controller - content for this file 
func GetAllMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// create movie 
func CreateMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// mark as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// delete movies 
func DeleteOneMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}