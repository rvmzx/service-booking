package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBStorage interface {
	Insert(context.Context) error
}

type MongoDB struct {
	client *mongo.Client
}

func SetupDatabase(ctx context.Context) (*MongoDB, error) {
	opts := options.Client().ApplyURI("mongodb+srv://serv:DdTQy5Ac3W5Hdb5@bookings.hktebqy.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	return &MongoDB{client: client}, nil
}

type MyDocument struct {
	ID    int    `bson:"_id,omitempty"`
	Name  string `bson:"name"`
	Value int    `bson:"value"`
}

func (m *MongoDB) Insert(ctx context.Context) error {
	collection := m.client.Database("bookings-service").Collection("bookings")

	document := MyDocument{
		Name:  "John Doe",
		Value: 42,
	}

	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return err
	}

	fmt.Println("Inserted document ID:", result.InsertedID)
	return nil
}
