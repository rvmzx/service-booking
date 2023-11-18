package api

import (
	"context"

	"github.com/rvmzx/service-booking/internal/storage"
)

type BookingManager struct {
	DBStorage storage.DBStorage
}

func NewBookingManager(db storage.DBStorage) *BookingManager {
	return &BookingManager{
		DBStorage: db,
	}
}

func (m *BookingManager) NewBooking(ctx context.Context, booking *storage.Booking) error {
	return m.DBStorage.Insert(ctx, booking)
}
