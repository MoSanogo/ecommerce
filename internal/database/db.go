package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

// DB represents the database connection pool
type DB struct {
	*sql.DB
}

// NewDB initializes a new database connection pool
func NewSQLiteDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)
	if err := db.Ping(); err != nil {
		db.Close()
		log.Printf("Failed to ping database: %v", err)
		// Return a wrapped error for better context
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{db}, nil
}

// InitSchema  creates the necessary tables in the database if they do not exist
func (db *DB) InitSchema() error {
	schemaSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,	
		username TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		roles TEXT NOT NULL,
		soft_delete BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS products (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		price REAL NOT NULL,
		stock_quantity INTEGER NOT NULL,
		image_urls TEXT,
		tenant_id TEXT NOT NULL,
		category TEXT NOT NULL,
		soft_delete BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (tenant_id) REFERENCES users(id)
	);
	CREATE TABLE IF NOT EXISTS carts (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL,
		product_id TEXT NOT NULL,
		quantity INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (product_id) REFERENCES products(id)
	);
	CREATE TABLE IF NOT EXISTS cart_items (
		id TEXT PRIMARY KEY,
		cart_id TEXT NOT NULL,
		product_id TEXT NOT NULL,
		quantity INTEGER NOT NULL,
		price REAL NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (cart_id) REFERENCES carts(id),
		FOREIGN KEY (product_id) REFERENCES products(id)
	);
	CREATE TABLE IF NOT EXISTS orders (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL,
		amount REAL NOT NULL,
		status TEXT NOT NULL,
		soft_delete BOOLEAN DEFAULT FALSE,
		shipping_address TEXT NOT NULL,
		shipping_status TEXT NOT NULL,
		shipping_tracking_number TEXT NOT NULL,
		order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	CREATE TABLE IF NOT EXISTS order_items (
		id TEXT PRIMARY KEY,
		order_id TEXT NOT NULL,
		product_id TEXT NOT NULL,
		quantity INTEGER NOT NULL,
		price REAL NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (order_id) REFERENCES orders(id),
		FOREIGN KEY (product_id) REFERENCES products(id)
	);
	CREATE TABLE IF NOT EXISTS payments (
		id TEXT PRIMARY KEY,
		order_id TEXT NOT NULL,
		amount REAL NOT NULL,
		status TEXT NOT NULL,
		payment_method TEXT NOT NULL,
		transaction_id TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (order_id) REFERENCES orders(id)
	);
	CREATE TABLE IF NOT EXISTS reviews (
		id TEXT PRIMARY KEY,
		product_id TEXT NOT NULL,
		user_id TEXT NOT NULL,
		rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
		comment TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (product_id) REFERENCES products(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	CREATE TABLE IF NOT EXISTS wishlists (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL,
		name TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (product_id) REFERENCES products(id)
	);
	CREATE TABLE IF NOT EXISTS wishlist_items (
		id TEXT PRIMARY KEY,
		wishlist_id TEXT NOT NULL,
		user_id TEXT NOT NULL,
		product_id TEXT NOT NULL,
		added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		removed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (wishlist_id) REFERENCES wishlists(id),
		FOREIGN KEY (product_id) REFERENCES products(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`
	_, err := db.Exec(schemaSQL)
	if err != nil {
		return fmt.Errorf("failed to initialize schema: %w", err)
	}

	log.Println("Database schema initialized successfully")
	return nil
}
func Cleanup(dataSourceName string) error {
	if dataSourceName == "" {
		return fmt.Errorf("data source name is empty")
	}

	os.Remove(dataSourceName)
	log.Printf("Database file %s removed successfully", dataSourceName)
	return nil
}
