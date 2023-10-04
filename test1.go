package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main2() {
	// Replace with your MongoDB Atlas connection string
	connectionString := "mongodb+srv://shailendrashail:mA6FwOa004ImXoEA@shailtestcluster.da31tjz.mongodb.net/?retryWrites=true&w=majority"

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB Atlas
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB Atlas!")

	// Disconnect when done
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// You can now perform database operations using the "client" object.
	// For example, you can create a collection and insert a document:
	collection := client.Database("shaildb").Collection("user")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert a document
	_, err = collection.InsertOne(ctx, bson.M{"author": "Shail"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Document inserted successfully!")
}
