package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kiloMIA/QuanMarketplace/backend/internal/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

func (ur *UserRepository) CreateUser (ctx context.Context, user *models.User) error {
	ur.lg.Debug("user repository level - CreateUser")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		ur.lg.Errorf("user repository level - GenerateFromPassword - error hashing plaintext password - %v", err)
		return err
	}

	query := "INSERT INTO users (name, email, password, enable) VALUES ($1, $2, $3, $4)"
	_, err = ur.DB.Exec(ctx, query, user.Name, user.Email, string(hashedPassword), user.Enable)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") && strings.Contains(err.Error(), "Email") {
			return fmt.Errorf("email %s already exists", user.Email)
		}
		ur.lg.Errorf("user repository level - CreateUser - db exec - %v", err)
		return err
	}

	return nil
}