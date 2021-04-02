package database

import (
	"fsp/open-music/packages/env"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Client {
	mongoPort := env.GetEnv("MONGO_PORT")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:" + mongoPort))
	if err != nil {
		log.Fatal(err)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:"+mongoPort))

	// ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()
	// err = client.Ping(ctx, readpref.Primary())

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	return client
}
