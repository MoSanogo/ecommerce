package models

import (
	"time"
)

type Payment struct {
	ID            string    `json:"id"`
	OrderID       string    `json:"order_id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	TransactionID string    `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
