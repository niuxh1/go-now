package jwt

import (
	"testing"
)

func TestJwtFlow(t *testing.T) {
	userId := int64(123)
	username := "testuser"
	role := "admin"

	// 1. Generate Token
	tokenString, err := GetToken(userId, username, role)
	if err != nil {
		t.Fatalf("GetToken failed: %v", err)
	}
	if tokenString == "" {
		t.Fatal("Generated token is empty")
	}

	// 2. Parse Token
	claims, err := ParseToken(tokenString)
	if err != nil {
		t.Fatalf("ParseToken failed: %v", err)
	}

	// 3. Verify Claims
	if claims.UserID != userId {
		t.Errorf("Expected UserID %d, got %d", userId, claims.UserID)
	}
	if claims.Username != username {
		t.Errorf("Expected Username %s, got %s", username, claims.Username)
	}
	if claims.Role != role {
		t.Errorf("Expected Role %s, got %s", role, claims.Role)
	}
}
