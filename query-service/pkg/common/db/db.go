package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	uri := os.Getenv("MONGO_URI")

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    "admin",
		Username:      "root",
		Password:      "root",
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetAuth(credential))

	if err != nil {
		panic(err)
	}

	return client
}
