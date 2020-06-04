package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main()  {
	clientOptions := options.Client().ApplyURI("몽고 DB 주소를 입력해주세요")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("몽고 DB에 연결했습니다!")
	uesrsCollection := client.Database("test").Collection("users")

	// find documents
	cursor, err := uesrsCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("에러!")
		fmt.Println(err)
	}
	for cursor.Next(context.TODO()) {
		var elem bson.M
		err := cursor.Decode(&elem)

		if err != nil {
			fmt.Println(err)
		}
		// find 결과 print
		fmt.Println(elem)
	}

	// create a document
	insertResult, _ := uesrsCollection.InsertOne(context.TODO(), bson.D{
        {"userID", "test1234"},
        {"array", bson.A{"flying", "squirrel", "dev"}},
	})
	fmt.Println(insertResult)

	// update a document
	updateFilter := bson.D{{"userID", "test1234"}}
	updateBson := bson.D{
		{"$push", bson.D{
			{"array", "wow"},
		}},
	}
	updateResult, _ := uesrsCollection.UpdateOne(
        context.TODO(),
        updateFilter,
        updateBson,
	)
	fmt.Println(updateResult)

	// delete a document
	deleteFilter := bson.D{{"userID", "test1234"}}
	deleteResult, _ := uesrsCollection.DeleteOne(
		context.TODO(),
		deleteFilter,
	)
	fmt.Println(deleteResult)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("몽고DB 연결을 종료했습니다!")
}