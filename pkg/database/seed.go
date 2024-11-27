package database

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// PopulateDB generates and inserts a large amount of synthetic data into the database.
func PopulateDB() {
	if Pool == nil {
		log.Fatal("Database connection pool is not initialized. Ensure you call InitDB() before running PopulateDB.")
	}

	// Fetch the last ID for each table
	var lastUserID, lastWarehouseID, lastItemID int
	// Fetch last user ID
	err := Pool.QueryRow(context.Background(), "SELECT last_value FROM users_id_seq").Scan(&lastUserID)
	if err != nil {
		log.Fatalf("Failed to fetch last user ID: %v", err)
	}
	// Fetch last warehouse ID
	err = Pool.QueryRow(context.Background(), "SELECT last_value FROM warehouses_warehouse_id_seq").Scan(&lastWarehouseID)
	if err != nil {
		log.Fatalf("Failed to fetch last warehouse ID: %v", err)
	}
	// Fetch last item ID
	err = Pool.QueryRow(context.Background(), "SELECT last_value FROM items_item_id_seq").Scan(&lastItemID)
	if err != nil {
		log.Fatalf("Failed to fetch last item ID: %v", err)
	}

	// Number of records to insert
	numUsers := 10
	numItems := 50
	numWarehouses := 5
	numWarehouseItems := 200

	// Insert users
	fmt.Println("Inserting users...")
	for i := 1; i <= numUsers; i++ {
		name := fmt.Sprintf("User%d", lastUserID+i)
		email := fmt.Sprintf("user%d@.com", lastUserID+i)
		password := fmt.Sprintf("%v", HashPassword(strconv.Itoa(lastUserID)))

		_, err := Pool.Exec(context.Background(),
			"INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
			name, email, password)
		if err != nil {
			log.Fatalf("Failed to insert user %d: %v", i, err)
		}
	}
	fmt.Printf("%d users inserted successfully.\n", numUsers)

	// Insert items
	fmt.Println("Inserting items...")
	for i := 1; i <= numItems; i++ {
		name := fmt.Sprintf("Item%d", lastItemID+i)
		category := fmt.Sprintf("Category%d", rand.Intn(10)+1) // 10 random categories
		description := fmt.Sprintf("Description for item %d", i)
		quantity := rand.Intn(100) + 1
		unitPrice := rand.Float64()*100 + 1 // Random price between 1 and 100
		totalPrice := float64(quantity) * unitPrice
		fmt.Println(lastUserID, lastItemID)
		_, err := Pool.Exec(context.Background(),
			"INSERT INTO items (name, category, description, quantity, unit_price, total_price) VALUES ($1, $2, $3, $4, $5, $6)",
			name, category, description, quantity, unitPrice, totalPrice)
		if err != nil {
			log.Fatalf("Failed to insert item %d: %v", i, err)
		}
	}
	fmt.Printf("%d items inserted successfully.\n", numItems)

	// Insert warehouses
	fmt.Println("Inserting warehouses...")
	for i := 1; i <= numWarehouses; i++ {
		location := fmt.Sprintf("Warehouse Location %d", lastWarehouseID+i)
		currentCapacity := rand.Intn(500) + 50            // Random capacity between 50 and 500
		totalCapacity := currentCapacity + rand.Intn(100) // Random total capacity > current capacity
		adminID := rand.Intn(numUsers+lastUserID) + 1     // Random admin ID from users

		_, err := Pool.Exec(context.Background(),
			"INSERT INTO warehouses (location, current_capacity, total_capacity, admin_id) VALUES ($1, $2, $3, $4)",
			location, currentCapacity, totalCapacity, adminID)
		if err != nil {
			log.Fatalf("Failed to insert warehouse %d: %v", i, err)
		}
	}
	fmt.Printf("%d warehouses inserted successfully.\n", numWarehouses)

	// Insert warehouse items
	fmt.Println("Inserting warehouse items...")
	for i := 1; i <= numWarehouseItems; i++ {
		warehouseID := rand.Intn(numWarehouses+lastWarehouseID) + 1
		itemID := rand.Intn(numItems+lastItemID) + 1
		quantity := rand.Intn(100) + 1

		_, err := Pool.Exec(context.Background(),
			"INSERT INTO warehouseItems (warehouse_id, item_id, quantity) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING",
			warehouseID, itemID, quantity)
		if err != nil {
			log.Fatalf("Failed to insert warehouse item %d: %v", i, err)
		}
	}
	fmt.Printf("%d warehouse items inserted successfully.\n", numWarehouseItems)

	fmt.Println("Database population complete.")
}

// Main function to execute the script
func main() {
	InitDB() // Initialize the database connection pool
	defer Pool.Close()

	startTime := time.Now()
	PopulateDB()
	fmt.Printf("Database populated in %v.\n", time.Since(startTime))
}
