package models

// This is the user details for authentication and authorization
// We could UserProfile to collect more info about user in our database.
type User struct {
	ID            string   `json:"id" validate:"required"`
	Username      string   `json:"username" validate:"required,min=10"`
	Email         string   `json:"email" validate:"required,email"`
	Password_hash string   `json:"_"`
	Roles         []string `json:"roles" validate:"required"` // e.g., ["admin", "user"]
}
