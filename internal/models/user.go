package models

type User struct {
	ID            string   `json:"id"`
	Username      string   `json:"username"`
	Email         string   `json:"email"`
	Password_hash string   `json:"_"`
	Roles         []string `json:"roles"` // e.g., ["admin", "user"]
}
