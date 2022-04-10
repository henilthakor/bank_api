package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/henilthakor/bank_api/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	//Number     string             `json:"number" bson:"number"`
	HolderName string  `json:"holdername" bson:"holdername"`
	Balance    float64 `json:"balance" bson:"balance"`
}

const uri = "mongodb://admin:admin@localhost:27017"
const dbname = "bankdb"
const colname = "accounts"

func AddAccount(a *Account) (interface{}, error) {
	client, ctx, cancel, err := helper.ConnectDB(uri)
	if err != nil {
		return nil, errors.New("Connection Error")
	}

	defer helper.CloseDB(client, ctx, cancel)

	collection := client.Database(dbname).Collection(colname)

	inserted, err := collection.InsertOne(context.Background(), a)
	if err != nil {
		return nil, errors.New("Insertion Error")
	}
	fmt.Println("Insert id of doc ", inserted.InsertedID)

	return inserted.InsertedID, nil
}

func GetAccountDetail(id string) (Account, error) {
	var account Account

	client, ctx, cancel, err := helper.ConnectDB(uri)
	if err != nil {
		return account, errors.New("Connection Error")
	}

	defer helper.CloseDB(client, ctx, cancel)

	collection := client.Database(dbname).Collection(colname)

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return account, errors.New("Query Error")
	}

	filter := bson.M{"_id": _id}

	var result bson.M

	if err := collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("Find result", result)

	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &account)

	return account, nil
}

func UpdateAccount(a *Account) error {
	client, ctx, cancel, err := helper.ConnectDB(uri)
	if err != nil {
		return errors.New("Connection Error")
	}

	defer helper.CloseDB(client, ctx, cancel)

	collection := client.Database(dbname).Collection(colname)

	filter := bson.M{"_id": a.ID}
	update := bson.M{"$set": bson.M{"name": a.HolderName, "balance": a.Balance}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("Query Error")
	}

	fmt.Println("Updated documents:", result.MatchedCount)
	return nil
}

func DeleteAccount(id string) error {
	client, ctx, cancel, err := helper.ConnectDB(uri)
	if err != nil {
		return errors.New("Connection Error")
	}

	defer helper.CloseDB(client, ctx, cancel)

	collection := client.Database(dbname).Collection(colname)

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("Account ID error")
	}

	filter := bson.M{"_id": _id}

	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return errors.New("Deletion Error")
	}

	fmt.Printf("Number of documents deleted: %d\n", result.DeletedCount)
	return nil
}
