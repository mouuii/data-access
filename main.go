package main

import (
	"context"
	"fmt"
	"github/mouuii/mongo/biz"
	"github/mouuii/mongo/data"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	mongo := data.NewMongo()
	d, _, e := data.NewData(mongo)
	if e != nil {
		log.Fatal(e)
	}
	cartR := data.NewCartRepo(d)
	c := &biz.Cart{
		UserId: 12,
		Items: []biz.Item{
			{Id: 1212, Quantity: 23},
			{Id: 1213, Quantity: 23},
		},
	}
	err := cartR.SaveCart(context.Background(), c)
	fmt.Println("xxx")
	if err != nil {
		fmt.Println("xxx %v", err)
	}
	// collectionTrains := client.Database("testing").Collection("trains")
	// insertOne(collectionTrains)
	// insertMany(collectionTrains)

	// collectionSaveCart := client.Database("testing").Collection("cart")
	// SaveCart(collectionSaveCart, ctx, &Trainer{})
	// test()

}

type Trainer struct {
	Name string
	Age  int
	City string
}

func insertOne(collection *mongo.Collection) {
	ash := Trainer{"Ash", 10, "Pallet Town"}
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func insertMany(collection *mongo.Collection) {
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}
	trainers := []interface{}{misty, brock}
	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}
