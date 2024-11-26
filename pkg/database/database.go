package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

// InitDB initializes the database connection pool.
func InitDB() *pgxpool.Pool {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), // Use DB_PASS instead of DB_PASSWORD
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Checking if connection string components are not empty
	if os.Getenv("DB_USER") == "" || os.Getenv("DB_PASS") == "" || os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_NAME") == "" {
		log.Fatalf("Database environment variables are not set correctly")
	}

	// Parsing the connection string into pgxpool config
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Unable to parse config: %v", err)
	}

	// Set connection pool configurations
	config.MaxConns = 10
	config.MaxConnLifetime = 30 * time.Minute

	// Create the connection pool without redeclaring err
	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(150) UNIQUE NOT NULL,
        password VARCHAR(100) NOT NULL
    );
    `

	// Execute the query
	_, err2 := Pool.Exec(context.Background(), query)
	if err2 != nil {
		log.Fatalf("Failed to create users table: %v", err2)
	} else {
		fmt.Println("Users table created or already exists.")
	}

	fmt.Println("Connected to PostgreSQL")
	return Pool
}
