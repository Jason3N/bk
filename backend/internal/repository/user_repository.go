package repository

import (
	"context"

	"github.com/Jason3n/bk/internal/models"
	"github.com/jackc/pgx"
)

// Define the interface
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
}

// Define the repository struct
type userRepository struct {
	conn *pgx.Conn
}

// Constructor for the repository
func NewUserRepository(conn *pgx.Conn) UserRepository {
	return &userRepository{conn: conn}
}

// Create a new user
func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO "user" (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	return r.conn.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
}
