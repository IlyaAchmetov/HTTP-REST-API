package model

import (
	"testing"
)

// TestUser ... создаем тестового пользователя
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.com",
		Password: "password",
	}
}
