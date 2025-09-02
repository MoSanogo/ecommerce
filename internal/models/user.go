package models

// This is the user details for authentication and authorization
// We could UserProfile to collect more info about user in our database.
type User struct {
	ID            string   `json:"id"`
	Username      string   `json:"username"`
	Email         string   `json:"email"`
	Password_hash string   `json:"_"`
	Roles         []string `json:"roles"` // e.g., ["admin", "user"]
}
