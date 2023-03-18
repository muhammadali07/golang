package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

func CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("mydatabase").Collection("users")
	result, err := collection.InsertOne(request.Context(), user)
	if err != nil {
		log.Fatal(err)
	}
	// log request and response
	log.Printf("Request: %s %s %s\n", request.Method, request.URL, request.Body)
	log.Printf("Response: %d %v\n", http.StatusCreated, result)

	json.NewEncoder(response).Encode(result)
}

func GetUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user User
	collection := client.Database("mydatabase").Collection("users")
	err := collection.FindOne(request.Context(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	// log request and response
	log.Printf("Request: %s %s\n", request.Method, request.URL)
	log.Printf("Response: %d %v\n", http.StatusOK, user)

	json.NewEncoder(response).Encode(user)
}

func GetAllUsersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var users []User
	collection := client.Database("mydatabase").Collection("users")
	cursor, err := collection.Find(request.Context(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(request.Context())
	for cursor.Next(request.Context()) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}

	// log request and response
	log.Printf("Request: %s %s\n", request.Method, request.URL)
	log.Printf("Response: %d %v\n", http.StatusOK, users)

	json.NewEncoder(response).Encode(users)
}

func UpdateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("mydatabase").Collection("users")
	result, err := collection.UpdateOne(request.Context(), bson.M{"_id": id}, bson.D{
		{"$set", bson.D{
			{"name", user.Name},
			{"email", user.Email},
			{"password", user.Password},
		}},
	})
	if err != nil {
		log.Fatal(err)
	}

	// log request and response
	log.Printf("Request: %s %s %s\n", request.Method, request.URL, request.Body)
	log.Printf("Response: %d %v\n", http.StatusOK, result)

	json.NewEncoder(response).Encode(result)
}

func DeleteUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("mydatabase").Collection("users")
	result, err := collection.DeleteOne(request.Context(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}

	// log request and response
	log.Printf("Request: %s %s\n", request.Method)

	json.NewEncoder(response).Encode(result)
}

func main() {
	ctx := context.Background()

	// set up client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// connect to MongoDB
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// create a new router
	router := mux.NewRouter()

	// define API endpoints
	router.HandleFunc("/users", CreateUserEndpoint).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", GetUserEndpoint).Methods(http.MethodGet)
	router.HandleFunc("/users", GetAllUsersEndpoint).Methods(http.MethodGet)
	router.HandleFunc("/users-update/{id}", UpdateUserEndpoint).Methods(http.MethodPut)
	router.HandleFunc("/users-delete/{id}", DeleteUserEndpoint).Methods(http.MethodDelete)

	// start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
