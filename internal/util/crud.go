package util

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Index show all the documents in a collection
func Index(Collection *mongo.Collection) []primitive.M {
	// retrieve all the documents in a collection
	cursor, err := Collection.Find(context.TODO(), bson.D{})
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}

// Show returns a single document in a collection
func Show(Collection *mongo.Collection, id string) []primitive.M {
	// retrieve all the documents in a collection
	cursor, err := Collection.Find(context.TODO(), bson.D{})
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}

// Create creates a new document in a collection
func Create(Collection *mongo.Collection, data interface{}) []primitive.M {
	// insert a single document into a collection
	// create a bson.D object
	result, err := Collection.InsertOne(context.TODO(), data)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)
	return Index(Collection)
}
