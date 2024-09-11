package repo

import (
	"ScaleSync/pkg/models"
	"database/sql"
	"errors"
)

type ItemRepository struct {
	DB *sql.DB
}

func (r *ItemRepository) Create(item *models.Item) error {
	query := `INSERT INTO items (name, category, description, quantity, unit_price, total_price) VALUES ($1, $2, $3, $4, $5, $6) RETURNING item_id`
	err := r.DB.QueryRow(query, item.Name, item.Category, item.Description, item.Quantity, item.UnitPrice, item.TotalPrice).Scan(&item.Item_ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ItemRepository) ReadAll() ([]*models.Item, error) {
	rows, err := r.DB.Query(`SELECT item_id, name, category, description, quantity, unit_price, total_price FROM items`)
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
	return items, nil
}

func (r *ItemRepository) Read(item_id int) (*models.Item, error) {
	var item models.Item
	err := r.DB.QueryRow(`SELECT item_id, name, category, description, quantity, unit_price, total_price FROM items WHERE item_id = $1`, item_id).
		Scan(&item.Item_ID, &item.Name, &item.Category, &item.Description, &item.Quantity, &item.UnitPrice, &item.TotalPrice)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("item not found")
		}
		return nil, err
	}

	return &item, nil
}

func (r *ItemRepository) Update(item *models.Item) error {
	query := `UPDATE items SET name = $1, category = $2, description = $3, quantity = $3, unit_price = $4, total_price = $5 WHERE id = $6`
	_, err := r.DB.Exec(query, item.Name, item.Category, item.Description, item.Quantity, item.UnitPrice, item.TotalPrice, item.Item_ID)
	return err
}

func (r *ItemRepository) Delete(item_id int) error {
	_, err := r.DB.Exec(`DELETE FROM items WHERE item_id = $1`, item_id)
	return err
}
