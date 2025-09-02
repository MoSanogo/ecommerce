package models

import (
	"time"
)

// Wishlist represents a user's wishlist
type Wishlist struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// WishlistItem represents an item in a wishlist
type WishlistItem struct {
	ID         string    `json:"id"`
	WishlistID string    `json:"wishlist_id"`
	ProductID  string    `json:"product_id"`
	UserID     string    `json:"user_id"`
	AddedAt    time.Time `json:"added_at"`
	RemovedAt  time.Time `json:"removed_at"`
}
