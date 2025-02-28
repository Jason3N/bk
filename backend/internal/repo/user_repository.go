package respoitory

import (
	"context"

	// own model
	"github.com/Jason3n/bk/internal/models"
	"github.com/jackc/pgx"
)

// define the interface for our access
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
}

// define the pgx connection using the repo interface
type userRespository struct {
	conn *pgx.Conn
}

/*
this will error out if you do not use the correct return statements
as defined in your interface
*/

// declare a new instance of user_repo
func NewUserRespository(conn *pgx.Conn) UserRepository {
	return &userRespository{conn: conn}
}

// create a new user
func (r *userRespository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO "user" (username, email, password) values ($1, $2, $3) RETURNING id`
	return r.conn.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
}
