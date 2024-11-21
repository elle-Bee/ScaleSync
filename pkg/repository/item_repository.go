package repository

import (
	"ScaleSync/pkg/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemRepository struct {
	DB *pgxpool.Pool
}

func (r *ItemRepository) Create(item *models.Item) error {
	query := `INSERT INTO items (name, category, description, quantity, unit_price, total_price) VALUES ($1, $2, $3, $4, $5, $6) RETURNING item_id`
	err := r.DB.QueryRow(context.Background(), query, item.Name, item.Category, item.Description, item.Quantity, item.UnitPrice, item.TotalPrice).Scan(&item.Item_ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ItemRepository) ReadAll() ([]*models.Item, error) {
	rows, err := r.DB.Query(context.Background(), `SELECT item_id, name, category, description, quantity, unit_price, total_price FROM items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.Item_ID, &item.Name, &item.Category, &item.Description, &item.Quantity, &item.UnitPrice, &item.TotalPrice); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) Read(itemID int) (*models.Item, error) {
	var item models.Item
	err := r.DB.QueryRow(context.Background(), `SELECT item_id, name, category, description, quantity, unit_price, total_price FROM items WHERE item_id = $1`, itemID).
		Scan(&item.Item_ID, &item.Name, &item.Category, &item.Description, &item.Quantity, &item.UnitPrice, &item.TotalPrice)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("item not found")
		}
		return nil, err
	}

	return &item, nil
}

// Get items by warehouse ID
func (r *ItemRepository) GetItemsByWarehouseID(warehouseID int) ([]models.Item, error) {
	var items []models.Item

	rows, err := r.DB.Query(context.Background(), `
		SELECT item_id, name, category, description, quantity, unit_price, total_price 
		FROM items 
		WHERE item_id IN (SELECT item_id
						FROM warehouseItems
						WHERE warehouse_id = $1)`, warehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.Item_ID, &item.Name, &item.Category, &item.Description, &item.Quantity, &item.UnitPrice, &item.TotalPrice); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) Update(item *models.Item) error {
	query := `UPDATE items SET name = $1, category = $2, description = $3, quantity = $4, unit_price = $5, total_price = $6 WHERE item_id = $7`
	_, err := r.DB.Exec(context.Background(), query, item.Name, item.Category, item.Description, item.Quantity, item.UnitPrice, item.TotalPrice, item.Item_ID)
	return err
}

func (r *ItemRepository) Delete(itemID int) error {
	_, err := r.DB.Exec(context.Background(), `DELETE FROM items WHERE item_id = $1`, itemID)
	return err
}
