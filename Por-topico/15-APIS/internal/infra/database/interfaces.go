package database

import (
	"github.com/osniantonio/goexpert/15-APIS/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
