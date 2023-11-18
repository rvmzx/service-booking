package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"customer_name,omitempty"`
	Value            float32            `bson:"value,omitempty"`
	Duration         time.Duration      `bson:"duration,omitempty"`
	ServiceLocations []string           `bson:"service_locations,omitempty"`
}

func (m *MongoDB) GetAllServices(ctx context.Context) ([]*Service, error) {
	collection := m.client.Database("bookings-service").Collection("services")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var services []*Service
	for cursor.Next(context.Background()) {
		var service Service
		if err := cursor.Decode(&service); err != nil {
			return nil, err
		}
		services = append(services, &service)
	}

	fmt.Printf("got %d service", len(services))

	return services, cursor.Err()
}

func (m *MongoDB) AddService(ctx context.Context, service *Service) error {
	collection := m.client.Database("bookings-service").Collection("services")

	result, err := collection.InsertOne(ctx, service)
	if err != nil {
		return err
	}

	fmt.Println("Inserted document ID:", result.InsertedID)
	return nil
}
