package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	DB *pgxpool.Pool
	lg *logrus.Logger
}

func NewUserRepository(db *pgxpool.Pool, lg *logrus.Logger) *UserRepository {
	return &UserRepository{
		DB: db,
		lg: lg,
	}
}