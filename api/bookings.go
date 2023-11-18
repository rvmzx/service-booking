package api

import (
	"context"

	"github.com/rvmzx/service-booking/internal/storage"
)

type BookingManager struct {
	dbStorage storage.DBStorage
}

func NewBookingManager(db storage.DBStorage) *BookingManager {
	return &BookingManager{
		dbStorage: db,
	}
}

func (m *BookingManager) NewBooking(ctx context.Context, booking *storage.Booking) error {
	return m.dbStorage.AddBooking(ctx, booking)
}

func (m *BookingManager) GetAllBookings(ctx context.Context) ([]*storage.Booking, error) {
	return m.dbStorage.GetAllBookings(ctx)
}
