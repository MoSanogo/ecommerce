package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims represents the JWT claims
type Claims struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for the given user ID and roles
func GenerateToken(userID string, roles []string) (string, error) {
	jwtSecret, Ok := os.LookupEnv("JWT_SECRET")
	if !Ok {
		return "", os.ErrNotExist // Return error if JWT_SECRET is not set
	}

	claims := Claims{
		UserID: userID,
		Roles:  roles,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "ecommerce-api",
			IssuedAt:  jwt.TimeFunc().Unix(),
			ExpiresAt: jwt.TimeFunc().Add(24 * time.Hour).Unix(), // Token valid for 24 hours
			Audience:  "ecommerce-users",
			NotBefore: jwt.TimeFunc().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// ValidateToken validates the JWT token and returns the claims if valid
func ValidateToken(tokenString string) (*Claims, error) {
	jwtSecret, Ok := os.LookupEnv("JWT_SECRET")
	if !Ok {
		return nil, os.ErrNotExist // Return error if JWT_SECRET is not set
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, jwt.ErrSignatureInvalid // Return error if token is invalid
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil // Return the claims if valid
	}

	return nil, jwt.ErrSignatureInvalid // Return error if claims are not of type Claims
}

// RefreshToken generates a new token with the same claims but a new expiration time
func RefreshToken(oldTokenString string) (string, error) {
	claims, err := ValidateToken(oldTokenString)
	if err != nil {
		return "", err // Return error if old token is invalid
	}

	// Update the expiration time
	claims.StandardClaims.ExpiresAt = jwt.TimeFunc().Add(24 * time.Hour).Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret, Ok := os.LookupEnv("JWT_SECRET")
	if !Ok {
		return "", os.ErrNotExist // Return error if JWT_SECRET is not set
	}

	return newToken.SignedString([]byte(jwtSecret))
}

// IsTokenExpired checks if the token is expired
func IsTokenExpired(tokenString string) (bool, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return false, err // Return error if token is invalid
	}

	return claims.StandardClaims.ExpiresAt < jwt.TimeFunc().Unix(), nil // Check if the token is expired
}
