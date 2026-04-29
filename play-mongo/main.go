package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongo")

	layout := "2006-01-02"

	e, _ := time.Parse(layout, time.Now().Format(layout))
	fmt.Println("end: ", e)
	coll := client.Database("testtime").Collection("time")
	_, errInsert := coll.InsertOne(ctx, bson.M{"date": e, "name": "fadly", "_id": "HM1"})
	if errInsert != nil {
		fmt.Errorf("error insert date :%s", errInsert)
	}

	var body struct {
		ID       string    `json:"_id"`
		Date     time.Time `json:"date"`
		Name     string    `json:"name"`
		MemberID string    `json:"memberID"`
	}

	yt, err := time.Parse(layout, time.Now().Format(layout))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("yt ", yt)

	hexByte, _ := hex.DecodeString("HM1")
	i, _ := primitive.ObjectIDFromHex(string(hexByte))
	fmt.Println("id: ", i)
	filter := bson.M{"_id": "HM1", "date": bson.M{"$gte": yt}}
	err = coll.FindOne(context.Background(), filter).Decode(&body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body.Name)

}
