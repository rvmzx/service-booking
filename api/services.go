package api

import (
	"context"

	"github.com/rvmzx/service-booking/internal/storage"
)

type ServiceManager struct {
	dbStorage storage.DBStorage
}

func NewServiceManager(db storage.DBStorage) *ServiceManager {
	return &ServiceManager{
		dbStorage: db,
	}
}

func (m *BookingManager) NewService(ctx context.Context, service *storage.Service) error {
	return m.dbStorage.AddService(ctx, service)
}

func (m *BookingManager) GetAllServices(ctx context.Context) ([]*storage.Service, error) {
	return m.dbStorage.GetAllServices(ctx)
}
