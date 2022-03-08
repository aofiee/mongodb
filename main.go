package main

import (
	"context"
	"log"
	"time"

	"github.com/aofiee/mongodb/handler"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Podcast struct {
		PodcastName string `bson:"podcast_name"`
	}
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://dd_hakka:password@localhost:27017/dd_hakka"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := client.Database("dd_hakka")
	log.Println(`Connected to db_hakka`, db)
	podcastsCollection := db.Collection("podcasts")
	log.Println(`Created collection podcasts`, podcastsCollection)

	podcast := Podcast{
		PodcastName: "The Go Programming Language",
	}
	log.Println(`Inserting podcast`, podcast)

	podcastResult, err := podcastsCollection.InsertOne(ctx, podcast)
	if err != nil {
		log.Println("Error inserting podcast: ", err)
	}
	log.Println("Inserted podcast with ID: ", podcastResult.InsertedID)
	//////
	app := fiber.New()

	app.Get("/graph", handler.PlaygroundHandler)
	app.Post("/query", handler.GraphqlHandler)
	app.Listen(":3000")
}
