package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	CustomerName    string             `bson:"customer_name,omitempty"`
	CustomerEmail   string             `bson:"customer_email,omitempty"`
	CustomerPhone   string             `bson:"customer_phone,omitempty"`
	Value           float32            `bson:"value,omitempty"`
	StartTime       time.Time          `bson:"start_time,omitempty"`
	Duration        time.Duration      `bson:"duration,omitempty"`
	ServiceName     string             `bson:"service_name,omitempty"`
	ServiceLocation string             `bson:"service_location,omitempty"`
}

func (m *MongoDB) GetAllBookings(ctx context.Context) ([]*Booking, error) {
	collection := m.client.Database("bookings-service").Collection("bookings")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var bookings []*Booking
	for cursor.Next(context.Background()) {
		var booking Booking
		if err := cursor.Decode(&booking); err != nil {
			return nil, err
		}
		bookings = append(bookings, &booking)
	}

	fmt.Printf("got %d bookings", len(bookings))

	return bookings, cursor.Err()
}

func (m *MongoDB) AddBooking(ctx context.Context, booking *Booking) error {
	collection := m.client.Database("bookings-service").Collection("bookings")

	result, err := collection.InsertOne(ctx, booking)
	if err != nil {
		return err
	}

	fmt.Println("Inserted document ID:", result.InsertedID)
	return nil
}
