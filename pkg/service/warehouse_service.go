package service

import (
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/repository"
	"errors"
)

type WarehouseService interface {
	CreateWarehouse(warehouse *models.Warehouse) error
	GetWarehouses() ([]*models.Warehouse, error)
	GetWarehouse(id int) (*models.Warehouse, error)
	UpdateWarehouse(warehouse *models.Warehouse) error
	DeleteWarehouse(id int) error
}

type WarehouseServiceImpl struct {
	Repo repository.WarehouseRepository
}

func (s *WarehouseServiceImpl) CreateWarehouse(warehouse *models.Warehouse) error {
	if warehouse.Warehouse_ID <= 0 {
		return errors.New("invalid warehouse data")
	}
	return s.Repo.Create(warehouse)
}

func (s *WarehouseServiceImpl) GetAll() ([]*models.Warehouse, error) {
	return s.Repo.GetAll()
}

func (s *WarehouseServiceImpl) GetByID(id int) (*models.Warehouse, error) {
	return s.Repo.GetByID(id)
}

func (s *WarehouseServiceImpl) UpdateWarehouse(warehouse *models.Warehouse) error {
	return s.Repo.Update(warehouse)
}

func (s *WarehouseServiceImpl) DeleteWarehouse(id int) error {
	return s.Repo.Delete(id)
}
