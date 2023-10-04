package test

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BenchmarkReadAllRows(b *testing.B) {
	// Replace with your MongoDB Atlas connection string
	connectionString := "mongodb+srv://shailendrashail:mA6FwOa004ImXoEA@shailtestcluster.da31tjz.mongodb.net/?retryWrites=true&w=majority"

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB Atlas
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		b.Fatalf("Error connecting to MongoDB Atlas: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Get a handle to the collection you want to read from
	collection := client.Database("shaildb").Collection("user")

	// Create a context with a timeout
	ctx := context.Background()

	// Define an empty filter to fetch all documents
	filter := bson.D{}

	for i := 0; i < b.N; i++ {
		// Find all documents in the collection
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			b.Fatalf("Error querying collection: %v", err)
		}
		defer cursor.Close(ctx)

		// Iterate through the result set
		for cursor.Next(ctx) {
			var result bson.M
			if err := cursor.Decode(&result); err != nil {
				b.Fatalf("Error decoding document: %v", err)
			}
		}

		// Check for errors from iterating over the cursor
		if err := cursor.Err(); err != nil {
			b.Fatalf("Error iterating over cursor: %v", err)
		}
	}
}
