package models

import (
	"time"
)

type Product struct {
	ID            string    `json:"id"`
	TenantID      string    `json:"tenant_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	StockQuantity int       `json:"stock_quantity"`
	Category      string    `json:"category"`
	ImageURLs     []string  `json:"image_urls"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
