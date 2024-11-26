package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"strings"

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
		os.Getenv("DB_PASS"),
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

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(150) UNIQUE NOT NULL,
			password VARCHAR(100) NOT NULL
		);
	`

	insertQuery := `
		INSERT INTO users (name, email, password) 
		VALUES 
		('aa', 'aa@example.com', 'aa'),
		('Bob Smith', 'bob.smith@example.com', 'p2'),
		('Charlie Brown', 'charlie.brown@example.com', 'p3'),
		('Diana Prince', 'diana.prince@example.com', 'p4'),
		('Eve Adams', 'eve.adams@example.com', 'p5');
	`
	// Execute CREATE TABLE query
	_, err = Pool.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Fatalf("Failed to execute create table query: %v", err)
	}

	// Process the INSERT query
	lines := strings.Split(insertQuery, "\n")
	var newValues []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "(") && strings.HasSuffix(line, "),") || strings.HasSuffix(line, ");") {
			line = strings.TrimSuffix(line, ",")
			line = strings.TrimSuffix(line, ");")
			line = strings.TrimPrefix(line, "(")
			line = strings.TrimSuffix(line, ")")
			parts := strings.Split(line, ", ")
			if len(parts) != 3 {
				log.Fatalf("Malformed value line: %v", line)
			}

			name := strings.Trim(parts[0], "'")
			email := strings.Trim(parts[1], "'")
			password := strings.Trim(parts[2], "'")

			hashedPassword := HashPassword(password)
			if err != nil {
				log.Fatalf("Failed to hash password: %v", err)
			}

			newValues = append(newValues, fmt.Sprintf("('%s', '%s', '%s')", name, email, hashedPassword))
		}
	}

	// Build new INSERT query
	finalInsertQuery := fmt.Sprintf(
		"INSERT INTO users (name, email, password) VALUES %s;",
		strings.Join(newValues, ","),
	)

	// Execute the new INSERT query
	rowCount := 0
	Pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM users").Scan(&rowCount)
	if rowCount == 0 {
		_, err = Pool.Exec(context.Background(), finalInsertQuery)
		if err != nil {
			log.Fatalf("Failed to execute insert query: %v", err)
		}

		fmt.Println("Users inserted successfully with hashed passwords.")
	}
	fmt.Println("Connected to PostgreSQL")
	return Pool
}
