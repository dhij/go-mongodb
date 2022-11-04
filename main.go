package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tasksDB *mongo.Database
var collection *mongo.Collection
var ctx = context.TODO()

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

func prepareDB() *mongo.Database {
	if tasksDB != nil {
		fmt.Println("DB already connected")
		return tasksDB
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	tasksDB = client.Database("tasker")

	return tasksDB
}

func main() {
	tasksDB := prepareDB()
	collection = tasksDB.Collection("tasks")

	task := &Task{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      "Learn Go",
		Completed: false,
	}

	createTask(task)
}

func createTask(task *Task) error {
	_, err := collection.InsertOne(ctx, task)
	return err
}
