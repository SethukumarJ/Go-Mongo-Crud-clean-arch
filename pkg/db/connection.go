package db

import (
	"fmt"
	"log"

	config "github.com/SethukumarJ/go-gin-clean-arch/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
)

func ConnectDatabase(cfg config.Config) (*mongo.Client, error) {

	ctx := context.TODO()

	//Opens database
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	// Create the "mongo_demo" database if it does not exist
	err = mongoclient.Database("mongo_demo").RunCommand(context.TODO(), bson.D{{"create", "users"}}).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to create 'mongo_demo' database: %v", err)
	}

	usercollection := mongoclient.Database("mongo_demo").Collection("users")
	fmt.Println(usercollection)

	log.Println("\nConnected to database:", "demo_db")

	return mongoclient, nil
}
