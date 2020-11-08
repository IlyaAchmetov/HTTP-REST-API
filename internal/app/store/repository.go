package store

import "github.com/IlyaAchmetov/HTTP-REST-API/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}