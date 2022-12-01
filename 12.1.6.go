package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func initDb() {
	co := options.Client().ApplyURI("mongodb://localhost:27017")
	c, err := mongo.Connect(context.TODO(), co)
	if err != nil {
		log.Fatal(err)
	}
	err2 := c.Ping(context.TODO(), nil)
	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println("连接成功！")
	}
	client = c
}

func insertMany() {
	c := client.Database("go_db").Collection("Student")
	ctx := context.TODO()
	id, _ := primitive.ObjectIDFromHex("6388a2b2d4d681f80d4cb6cb")
	uqdate := bson.D{{"$set", bson.D{{"name", "big tom"}, {"age", 22}}}}
	ur, err := c.UpdateMany(ctx, bson.D{{"_id", id}}, uqdate)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("ur: %v\n", ur.ModifiedCount)
	}
}

func main() {
	initDb()
	insertMany()
}
