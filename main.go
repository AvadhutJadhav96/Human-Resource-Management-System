package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-here"
const mongoURI = "mongodb://localhost:27017"

type Employee struct {
	ID     string  `json:"id,omitempty" bson: "_id,omniempty`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`

	//mongo responses in bson format and from server we get json format
	//so except id everything is same in bson and json
	//its more like we are telling ID as string in Go is represented as id in json and _id in bson
	//generally we need to do marshalling and unmarshalling but fiber makes it more easy....
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
		query := bson.D{{}} //empty query to get all records from database

		var employees []Employee = make([]Employee, 0)
		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		//cursor points to all employees in the database

		if err := cursor.All(c.Context(), &employees); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		//here we are converting the content of cursor which is bson format to format go can understand

		return c.JSON(employees)
		//again converting it back to JSON so as response to say postman since it can understand only json format

	})
	//here c is for response and request



	app.Post("/employee", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")

		var employee Employee

		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		//converts the post request from postman into struct format which go understands

		//now we have converted the reqeust into our struct format
		//we will set the ID value to empty, this is way to tell mongo db that while saving the data
		//of employee create your own id

		employee.ID = ""

		insertionResult, err := collection.InsertOne(c.Context(), employee)
		//we saved it in database and saved the result in insertionResult

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		//now to confirm if the data is added in database lets return it back

		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
		//query to get the id

		createdRecord := collection.FindOne(c.Context(), filter)
		//more like a reverse engineering to get back the record we just saved

		var createdEmployee Employee

		createdRecord.Decode(createdEmployee)
		//decoding it in struct format the record we got from mongo

		return c.Status(201).JSON(createdEmployee)
		//returned to frontend back in JSON format
	})

	
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
}
