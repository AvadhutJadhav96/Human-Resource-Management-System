package main

import (
	"log"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoInstance struct {
	Client	*mongo.Client
	Db		*mongo.Database
}

var mg MongoInstance

const dbName = "fiber-here"
const mongoURI = "mongodb://localhost:27017"

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    float64
}

func Connect() error {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Connect to MongoDB directly using mongo.Connect()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	// Verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	// Assign database instance
	db := client.Database(dbName)

	// Store MongoDB instance globally
	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New() //just like starting an app in express

	app.Get("/employee", func(c *fiber.Ctx) error {

	})
	//here c is for response and request

	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
}
