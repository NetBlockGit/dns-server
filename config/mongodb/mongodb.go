package mongodb

import (
	"context"
	"dnsserver/util/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var StatsCollection *mongo.Collection

var client *mongo.Client

func Init() {
	uri := env.Get("MONGO_DB_URI")
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	db := client.Database("dnsBlockerDatabase")
	StatsCollection = db.Collection("stats")
}

func Disconnect() {
	client.Disconnect(context.Background())
}
