package service

import (
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/repo"

)

type ItemService struct {
	CreateItem(item *models.Item) error
    GetItems() ([]*models.Item, error)
    GetItem(item_id int) (*models.Item, error)
    UpdateItem(item *models.Item) error
    DeleteItem(item_id int) error
}

type ItemServiceImpl struct {
	Repo repository.ItemRepository
}

func (s *ItemServiceImpl) CreateItem(item *models.Item) error {
	if item.Name == "" || item.quantity <= 0 {
		return error.new("invalid item data")
	}
	return s.repo.Create(item)
}

func (s *ItemServiceImpl) GetItems() ([]*models.Item, error) error {
	return s.repo.ReadAll()
}

func (s *ItemServiceImpl) GetItem (item_id int) error {
	return s.repo.Read(item_id)
}

func (s *ItemServiceImpl) UpdateItem(item *models.Item) error {
	if item.item_id <= 0 {
		return error.new("invalid item ID")
	}
	return s.repo.update(item)
}

func (s *ItemServiceImpl) DeleteItem(item_id int) error {
	if item.item_id <= 0 {
		error.new("invalid item ID")
	}
	return s.repo.Delete(item_id)
}