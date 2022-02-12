package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var StatsCollection *mongo.Collection

func Init() {
	uri := os.Getenv("MONGO_DB_URI")
	var err error
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	db := client.Database("dnsBlockerDatabase")
	StatsCollection = db.Collection("stats")
}
