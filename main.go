package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type MongoInstance struct {
	Client
	Db
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
