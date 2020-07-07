package db

import (
	"context"

	"github.com/wolf00/golang_lms/user/constants"

	log "github.com/micro/go-micro/v2/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo blah
func ConnectMongo(ctx context.Context) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(constants.MongoConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to MongoDB!...")

	return client.Database(constants.DatabaseName)
}

// Users is used to retrieve collection from database
func Users(ctx context.Context) *mongo.Collection {
	return ConnectMongo(ctx).Collection(constants.Users)
}

// Sources is used to retrieve collection from database
func Sources(ctx context.Context) *mongo.Collection {
	return ConnectMongo(ctx).Collection(constants.Sources)
}

// Organizations is used to retrieve collection from database
func Organizations(ctx context.Context) *mongo.Collection {
	return ConnectMongo(ctx).Collection(constants.Organizations)
}
