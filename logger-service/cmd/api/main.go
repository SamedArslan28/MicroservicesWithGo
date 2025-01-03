package main

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://localhost:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
}

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		panic(err)
	}
	client = mongoClient

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	connect, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}
	return connect, nil
}
