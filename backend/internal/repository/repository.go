package repository

import (
	"context"

	"github.com/kiloMIA/QuanMarketplace/backend/internal/models"
)

type Repository struct {
	User
}

type User interface {
	CreateUser (ctx context.Context, user *models.User) error
}