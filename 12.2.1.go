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

func del() {
	c := client.Database("go_db").Collection("Student")
	ctx := context.TODO()
	id, _ := primitive.ObjectIDFromHex("6388a2b2d4d681f80d4cb6cb")
	dr, err := c.DeleteMany(ctx, bson.D{{"_id", id}})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("dr: %v\n", dr.DeletedCount)
	}
}

func main() {
	initDb()
	del()
}
