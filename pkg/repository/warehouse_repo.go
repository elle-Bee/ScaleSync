package repository

import (
	"ScaleSync/pkg/models"
	"database/sql"
	"errors"
)

type WarehouseRepository struct {
	DB *sql.DB
}

func (r *WarehouseRepository) Create(warehouse *models.Warehouse) error {
	query := `INSERT INTO warehouses (location, current_capacity, total_capacity, admin_id) VALUES ($1, $2, $3, $4) RETURNING warehouse_id`
	err := r.DB.QueryRow(query, warehouse.Location, warehouse.CurrentCapacity, warehouse.TotalCapacity, warehouse.Admin.ID).Scan(&warehouse.Warehouse_ID)
	if err != nil {
		return err
	}

	for _, item := range warehouse.Items {
		_, err := r.DB.Exec(`INSERT INTO warehouse_items (warehouse_id, item_id) VALUES ($1, $2)`, warehouse.Warehouse_ID, item.Item_ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *WarehouseRepository) GetAll() ([]*models.Warehouse, error) {
	rows, err := r.DB.Query(`SELECT warehouse_id, location, current_capacity, total_capacity, admin_id FROM warehouses`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []*models.Warehouse
	for rows.Next() {
		var warehouse models.Warehouse
		var adminID int
		if err := rows.Scan(&warehouse.Warehouse_ID, &warehouse.Location, &warehouse.CurrentCapacity, &warehouse.TotalCapacity, &adminID); err != nil {
			return nil, err
		}
		warehouse.Admin.ID = adminID // Assuming user information is handled elsewhere
		warehouses = append(warehouses, &warehouse)
	}
	return warehouses, nil
}

func (r *WarehouseRepository) GetByID(id int) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	err := r.DB.QueryRow(`SELECT warehouse_id, location, current_capacity, total_capacity, admin_id FROM warehouses WHERE warehouse_id = $1`, id).
		Scan(&warehouse.Warehouse_ID, &warehouse.Location, &warehouse.CurrentCapacity, &warehouse.TotalCapacity, &warehouse.Admin.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("warehouse not found")
		}
		return nil, err
	}
	// Fetch items for the warehouse
	rows, err := r.DB.Query(`SELECT item_id FROM warehouse_items WHERE warehouse_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var itemID int
		if err := rows.Scan(&itemID); err != nil {
			return nil, err
		}
		// Fetch item details by ID
		var item models.Item
		err := r.DB.QueryRow(`SELECT item_id, name, category, description, quantity, unit_price, total_price FROM items WHERE id = $1`, itemID).
			Scan(&item.Item_ID, &item.Name, &item.Description, &item.Quantity, &item.UnitPrice, &item.TotalPrice)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	warehouse.Items = items

	return &warehouse, nil
}

func (r *WarehouseRepository) Update(warehouse *models.Warehouse) error {
	query := `UPDATE warehouses SET location = $1, current_capacity = $2, total_capacity = $3, admin_id = $4 WHERE warehouse_id = $5`
	_, err := r.DB.Exec(query, warehouse.Location, warehouse.CurrentCapacity, warehouse.TotalCapacity, warehouse.Admin.ID, warehouse.Warehouse_ID)
	if err != nil {
		return err
	}
	// Handle item relationships
	// Remove existing items and add new ones
	_, err = r.DB.Exec(`DELETE FROM warehouse_items WHERE warehouse_id = $1`, warehouse.Warehouse_ID)
	if err != nil {
		return err
	}
	for _, item := range warehouse.Items {
		_, err := r.DB.Exec(`INSERT INTO warehouse_items (warehouse_id, item_id) VALUES ($1, $2)`, warehouse.Warehouse_ID, item.Item_ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *WarehouseRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM warehouses WHERE warehouse_id = $1`, id)
	return err
}
