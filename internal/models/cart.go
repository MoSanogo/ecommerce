package models

import (
	"time"
)

type Cart struct {
	ID         string     `json:"id"`
	TenantID   string     `json:"tenant_id"`
	UserID     string     `json:"user_id"`
	Items      []CartItem `json:"items"`
	TotalPrice float64    `json:"total_price"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
