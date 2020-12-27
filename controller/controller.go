package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pawanpaudel93/go-mux-restapi/helper"
	"github.com/pawanpaudel93/go-mux-restapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var collection *mongo.Collection

func InitDatabase() {
	client = helper.ConnectDB()
	collection = client.Database(os.Getenv("DATABASE_NAME")).Collection("books")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []models.Book

	// collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("books")
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)
	}

	if err := cursor.Err(); err != nil {
		helper.GetError(err, w)
	}
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	params := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	// collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("books")
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	// collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("books")
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var book models.Book

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&book)

	update := bson.D{
		{"$set", bson.D{
			{"isbn", book.ISBN},
			{"title", book.Title},
			{"author", bson.D{
				{"firstname", book.Author.FirstName},
				{"lastname", book.Author.LastName},
			}},
		}},
	}
	// collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("books")
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	book.ID = id
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	// collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("books")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}
