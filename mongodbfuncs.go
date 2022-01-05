package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type UserJSON struct {
	EmailAddress string `bson:"email_address,omitempty"`
}

var (
	mongodb_users = &mongo.Collection{}
)

func initializeMongodb() {
	uri := "mongodb://127.0.0.1:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	mongodb_users = client.Database("heap_overflow").Collection("users")
}

func addUserToDatabase(email string) error {
	newUser := UserJSON{email}
	_, err := mongodb_users.InsertOne(context.TODO(), newUser)
	return err
}
