package mongodb

import (
	"context"
	"fmt"
	"log"

	"019.restful_mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBNAME Database name
const DBNAME = "phonebook"

// COLLECTION Collection name
const COLLECTION = "users"

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

var db *mongo.Database

// Connect establish a connection to database
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}

func GetUsers() []models.User {
	cur, err := db.Collection(COLLECTION).Find(context.Background(), bson.D{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var elements []models.User
	var elem models.User
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements
}

func GetUser(id string) models.User {
	filter := bson.D{{"id", id}}
	var user models.User
	_ = db.Collection(COLLECTION).FindOne(context.Background(), filter).Decode(&user)
	return user
}

func CreateUser(user models.User) {
	fmt.Println(user)
	_, err := db.Collection(COLLECTION).InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUser(user models.User, id string) {
	filter := bson.D{{"id", id}}

	update := bson.D{
		{"$set", bson.D{
			{"firstname", user.Firstname},
			{"lastname", user.Lastname},
			{"info.city", user.City},
			{"info.phone", user.Phone},
		}},
	}
	_, err := db.Collection(COLLECTION).UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUser(id string) {
	_, err := db.Collection(COLLECTION).DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
}
