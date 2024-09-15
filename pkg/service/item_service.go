package service

import (
	"errors"
	"learn/pkg/models"
	"learn/pkg/repository"
)

type ItemService interface {
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
	if item.Name == "" || item.Quantity <= 0 {
		return errors.New("invalid item data")
	}
	return s.Repo.Create(item)
}

func (s *ItemServiceImpl) GetItems() ([]*models.Item, error) {
	return s.Repo.ReadAll()
}

func (s *ItemServiceImpl) GetItem(item_id int) (*models.Item, error) {
	return s.Repo.Read(item_id)
}

func (s *ItemServiceImpl) UpdateItem(item *models.Item) error {
	if item.Item_ID <= 0 {
		return errors.New("invalid item ID")
	}
	return s.Repo.Update(item)
}

func (s *ItemServiceImpl) DeleteItem(item_id int) error {
	if item_id <= 0 {
		return errors.New("invalid item ID")
	}
	return s.Repo.Delete(item_id)
}
