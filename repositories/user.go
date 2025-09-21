package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tumeraltunbas/go-blog/database/queries"
	"github.com/tumeraltunbas/go-blog/models/entities"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := queries.GetUserByEmailQuery(email)

	var user entities.User

	err := r.pool.QueryRow(ctx, query).Scan(
		&user.Id, &user.Email, &user.CreatedAt,
	)

	return &user, err
}

func (r *UserRepository) InsertUser(ctx context.Context, email string, password string) error {
	query := queries.InsertUser(email, password)
	_, err := r.pool.Exec(ctx, query)
	return err
}
