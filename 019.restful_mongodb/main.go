package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"019.restful_mongodb/handlers"
	"019.restful_mongodb/models"

	"github.com/gorilla/mux"
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

func init() {
	var users []models.User

	// 指定連線位置
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
	if err != nil {
		log.Fatal(err)
	}

	// 連線
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 建立 db
	db := client.Database(DBNAME)

	// 讀取假資料
	byteValues, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	// 轉換 data 成 struct
	json.Unmarshal(byteValues, &users)

	// 假資料塞到 userList slice
	var userList []interface{}
	for _, u := range users {
		userList = append(userList, u)
	}

	// 清空db
	_, err = db.Collection(COLLECTION).DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	// 假資料匯入 db
	_, err = db.Collection(COLLECTION).InsertMany(context.Background(), userList)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	fmt.Println("Starting server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
