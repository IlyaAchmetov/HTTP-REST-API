package model

import (
	"testing"
)

// TestUser ...https://youtu.be/vK8UY9fqLSY?t=277
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.com",
		Password: "password",
	}
}
