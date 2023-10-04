package test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func performReadOperation(collection *mongo.Collection, ctx context.Context, wg *sync.WaitGroup, b *testing.B, c int) {
	defer wg.Done()
	filter := bson.D{}
	fmt.Println(fmt.Sprintf("Start C -%d", c))
	for i := 0; i < b.N; i++ {
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			b.Fatalf("Error querying collection: %v", err)
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var result bson.M
			if err := cursor.Decode(&result); err != nil {
				b.Fatalf("Error decoding document: %v", err)
			}
		}
		if err := cursor.Err(); err != nil {
			b.Fatalf("Error iterating over cursor: %v", err)
		}
	}
	fmt.Println(fmt.Sprintf("Finish C -%d", c))
}

func BenchmarkConcurrentReads(b *testing.B) {
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

	// Number of concurrent goroutines
	numGoroutines := 1000

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launch concurrent read operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go performReadOperation(collection, ctx, &wg, b, i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
