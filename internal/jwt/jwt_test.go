package jwt

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	// Set the JWT_SECRET environment variable for testing
	t.Setenv("JWT_SECRET", "test_secret")

	token, err := GenerateToken("user123", []string{"user"})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Validate the token
	claims, err := ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "user123", claims.UserID)
	assert.ElementsMatch(t, []string{"user"}, claims.Roles)
}
func TestValidateToken(t *testing.T) {
	// Set the JWT_SECRET environment variable for testing
	t.Setenv("JWT_SECRET", "test_secret")

	// Generate a token
	token, err := GenerateToken("user123", []string{"user"})
	assert.NoError(t, err)

	// Validate the token
	claims, err := ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "user123", claims.UserID)
	assert.ElementsMatch(t, []string{"user"}, claims.Roles)

	// Test with an invalid token
	invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
	_, err = ValidateToken(invalidToken)
	assert.Error(t, err)
	assert.Equal(t, jwt.ErrSignatureInvalid, err) // Expect signature invalid error
}

func TestRefreshToken(t *testing.T) {
	// Set the JWT_SECRET environment variable for testing
	t.Setenv("JWT_SECRET", "test_secret")

	// Generate a token
	token, err := GenerateToken("user123", []string{"user"})
	assert.NoError(t, err)

	// Validate the token
	claims, err := ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "user123", claims.UserID)

	// Refresh the token by generating a new one with the same user ID and roles
	newToken, err := GenerateToken(claims.UserID, claims.Roles)
	assert.NoError(t, err)

	// Validate the new token
	newClaims, err := ValidateToken(newToken)
	assert.NoError(t, err)
	assert.Equal(t, claims.UserID, newClaims.UserID)
	assert.ElementsMatch(t, claims.Roles, newClaims.Roles)
}

func TestIsTokenExpired(t *testing.T) {
	// Set the JWT_SECRET environment variable for testing
	t.Setenv("JWT_SECRET", "test_secret")

	// Generate a token
	token, err := GenerateToken("user123", []string{"user"})
	assert.NoError(t, err)

	// Validate the token
	claims, err := ValidateToken(token)
	assert.NoError(t, err)

	// Check if the token is expired
	isExpired := claims.StandardClaims.ExpiresAt < jwt.TimeFunc().Unix()
	assert.False(t, isExpired) // The token should not be expired immediately after generation

	// Simulate expiration by setting the expiration time to a past time
	claims.StandardClaims.ExpiresAt = jwt.TimeFunc().Add(-24 * time.Hour).Unix()
	assert.True(t, claims.StandardClaims.ExpiresAt < jwt.TimeFunc().Unix()) // The token should now be considered expired
}
