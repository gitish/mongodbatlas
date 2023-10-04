package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main3() {
	// Replace with your MongoDB Atlas connection string
	connectionString := "mongodb+srv://shailendrashail:mA6FwOa004ImXoEA@shailtestcluster.da31tjz.mongodb.net/?retryWrites=true&w=majority"

	// Set client options with connection pooling
	clientOptions := options.Client().
		ApplyURI(connectionString).
		SetMaxPoolSize(100) // Adjust as needed

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

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Number of goroutines for concurrent read and write operations
	numGoroutines := 10 // Adjust as needed

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			performWriteOperations(client, id)
		}(i)
	}

	performReadOperations(client)

	// Wait for all goroutines to finish
	wg.Wait()
}

func performWriteOperations(client *mongo.Client, id int) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get a handle to the MongoDB collection
	collection := client.Database("shaildb").Collection("user")

	// Insert a document
	_, err := collection.InsertOne(ctx, bson.M{"user": fmt.Sprintf("user%d", id), "pass": fmt.Sprintf("pass%d", id)})
	if err != nil {
		log.Printf("Error inserting document: %v", err)
	} else {
		fmt.Printf("Document inserted by goroutine %d\n", id)
	}

}

func performReadOperations(client *mongo.Client) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get a handle to the MongoDB collection
	collection := client.Database("shaildb").Collection("user")

	// Define an empty filter to fetch all documents
	filter := bson.D{}

	// Find all documents in the collection
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	// Iterate through the result set
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
