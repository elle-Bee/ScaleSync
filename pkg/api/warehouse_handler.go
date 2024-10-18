package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"ScaleSync/pkg/database"
	"ScaleSync/pkg/models"

	"github.com/gorilla/mux"
)

// CreateWarehouse handles the creation of a new warehouse
func CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouse models.Warehouse
	err := json.NewDecoder(r.Body).Decode(&warehouse)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if warehouse with the same ID already exists
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_id = $1)`
	err = database.Pool.QueryRow(context.Background(), checkQuery, warehouse.Warehouse_ID).Scan(&exists)
	if err != nil {
		http.Error(w, "Error checking warehouse existence", http.StatusInternalServerError)
		log.Println("Check Warehouse Error:", err)
		return
	}
	if exists {
		http.Error(w, "Warehouse already exists", http.StatusBadRequest)
		return
	}

	// Insert new warehouse into the database
	query := `INSERT INTO warehouses (warehouse_id, location, current_capacity, total_capacity) 
	          VALUES ($1, $2, $3, $4) RETURNING warehouse_id`
	err = database.Pool.QueryRow(context.Background(), query, warehouse.Warehouse_ID, warehouse.Location, 0, warehouse.TotalCapacity).Scan(&warehouse.Warehouse_ID)
	if err != nil {
		http.Error(w, "Error creating warehouse", http.StatusInternalServerError)
		log.Println("Create Warehouse Error:", err)
		return
	}

	json.NewEncoder(w).Encode(warehouse)
}

// GetWarehouse handles fetching a single warehouse by ID
func GetWarehouse(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var warehouse models.Warehouse
	query := `SELECT warehouse_id, location, current_capacity, total_capacity FROM warehouses WHERE warehouse_id = $1`
	err := database.Pool.QueryRow(context.Background(), query, id).Scan(&warehouse.Warehouse_ID, &warehouse.Location, &warehouse.CurrentCapacity, &warehouse.TotalCapacity)
	if err != nil {
		http.Error(w, "Warehouse not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(warehouse)
}

// UpdateWarehouse handles updating an existing warehouse
func UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var warehouse models.Warehouse
	err := json.NewDecoder(r.Body).Decode(&warehouse)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := `UPDATE warehouses SET location=$1, current_capacity=$2, total_capacity=$3 WHERE warehouse_id=$4`
	_, err = database.Pool.Exec(context.Background(), query, warehouse.Location, warehouse.CurrentCapacity, warehouse.TotalCapacity, id)
	if err != nil {
		http.Error(w, "Error updating warehouse", http.StatusInternalServerError)
		log.Println("Update Warehouse Error:", err)
		return
	}

	json.NewEncoder(w).Encode(warehouse)
}

// DeleteWarehouse handles the deletion of a warehouse by ID
func DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	query := `DELETE FROM warehouses WHERE warehouse_id=$1`

	_, err := database.Pool.Exec(context.Background(), query, id)
	if err != nil {
		http.Error(w, "Error deleting warehouse", http.StatusInternalServerError)
		log.Println("Delete Warehouse Error:", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllWarehouses handles fetching all warehouses
func GetAllWarehouses(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Pool.Query(context.Background(), `SELECT warehouse_id, location, current_capacity, total_capacity FROM warehouses`)
	if err != nil {
		http.Error(w, "Error fetching warehouses", http.StatusInternalServerError)
		log.Println("Get All Warehouses Error:", err)
		return
	}
	defer rows.Close()

	var warehouses []models.Warehouse
	for rows.Next() {
		var warehouse models.Warehouse
		err = rows.Scan(&warehouse.Warehouse_ID, &warehouse.Location, &warehouse.CurrentCapacity, &warehouse.TotalCapacity)
		if err != nil {
			log.Println("Error scanning warehouse:", err)
			continue
		}
		warehouses = append(warehouses, warehouse)
	}

	json.NewEncoder(w).Encode(warehouses)
}
