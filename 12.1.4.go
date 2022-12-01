package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Student struct {
	Name string
	Age  int
}

func insert() {
	s1 := Student{
		Name: "tom",
		Age:  20,
	}
	c := client.Database("go_db").Collection("Student")
	ior, err := c.InsertOne(context.TODO(), s1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("ior.InsertedID: %v\n", ior.InsertedID)
	}
}

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
	s1 := Student{
		Name: "kiite",
		Age:  20,
	}
	s2 := Student{
		Name: "rose",
		Age:  20,
	}
	stus := []interface{}{s1, s2}
	imr, err := c.InsertMany(context.TODO(), stus)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("imr: %v\n", imr.InsertedIDs...)
	}

}

func mybson() {
	d := bson.D{{"name", "tom"}}
	fmt.Printf("d: %v\n", d)
}

func main() {
	initDb()
	//insertMany()
	//insert()
}
