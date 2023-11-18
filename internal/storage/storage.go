package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBStorage interface {
	GetAllBookings(context.Context) ([]*Booking, error)
	AddBooking(context.Context, *Booking) error

	GetAllServices(context.Context) ([]*Service, error)
	AddService(context.Context, *Service) error
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
