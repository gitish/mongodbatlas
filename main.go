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

func main() {
	client, error := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://shailendrashail:mA6FwOa004ImXoEA@shailtestcluster.da31tjz.mongodb.net/?retryWrites=true&w=majority"))
	if error != nil {
		log.Fatal(error)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	error = client.Connect(ctx)
	if error != nil {
		fmt.Println("Error: 0")
		log.Fatal(error)
	}
	defer client.Disconnect(ctx)

	database, error := client.ListDatabaseNames(ctx, bson.M{})
	if error != nil {
		fmt.Println("Error: 1ß")
		log.Fatal(error)
	}
	fmt.Println(database)

}
