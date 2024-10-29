package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:rblvFnBAXeSDGF86@cluster0.hodkn.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	MongoClient = client
	return client, nil
}

func GetMongoCollection(database, collection string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}
