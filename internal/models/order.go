package models

import (
	"time"
)

type Order struct {
	ID                     string      `json:"id"`
	UserID                 string      `json:"user_id"`
	Status                 string      `json:"status"`
	Amount                 float64     `json:"amount"`
	ShippingAddress        string      `json:"shipping_address"`
	ShippingStatus         string      `json:"shipping_status"`
	ShippingTrackingNumber string      `json:"shipping_tracking_number"`
	OrderDate              time.Time   `json:"order_date"`
	Items                  []OrderItem `json:"items"`
	CreatedAt              time.Time   `json:"created_at"`
	UpdatedAt              time.Time   `json:"updated_at"`
}
type OrderItem struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type OrderStatus string
