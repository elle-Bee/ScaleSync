package service

import "ScaleSync/pkg/models"

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