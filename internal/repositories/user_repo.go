package repositories

import (
	"context"
	entity "easy-ride-api/internal/models"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrInvalidCredentials = errors.New("Invalid email or password")
)

// Interface — defines the contract. Other layers depend on this, not the concrete type.
type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	CreateSession(ctx context.Context, user *entity.Session) (*entity.Session, error)
	InvalidateSession(ctx context.Context, token string, expiresAt time.Time) error
}

// Concrete implementation — holds the connection pool.
type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepository{pool: pool}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `
		INSERT INTO users (full_name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, full_name, email, is_active, created_at, updated_at`

	created := &entity.User{}
	err := r.pool.QueryRow(ctx, query,
		user.FullName,
		user.Email,
		user.Password,
	).Scan(
		&created.ID,
		&created.FullName,
		&created.Email,
		&created.IsActive,
		&created.CreatedAt,
		&created.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, full_name, email, password, is_active, created_at, updated_at FROM users WHERE email = $1`

	user := &entity.User{}
	err := r.pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {

			return nil, ErrInvalidCredentials
		}

		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	query := `SELECT id, full_name, email, is_active, created_at, updated_at FROM users WHERE id = $1`

	user := &entity.User{}
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) CreateSession(ctx context.Context, data *entity.Session) (*entity.Session, error) {
	query := `
	INSERT INTO sessions (user_id, token, expires_at, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, user_id, token, expires_at, created_at, updated_at
	`

	session := &entity.Session{}

	err := r.pool.QueryRow(ctx, query, data.UserId, data.Token, data.ExpiresAt, data.CreatedAt, data.UpdatedAt).Scan(
		&session.ID,
		&session.UserId,
		&session.Token,
		&session.ExpiresAt,
		&session.CreatedAt,
		&session.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *userRepository) InvalidateSession(ctx context.Context, token string, expiresAt time.Time) error {

	query := `UPDATE sessions SET expires_at = $1 WHERE token = $2;`

	_, err := r.pool.Exec(ctx, query, expiresAt, token)

	if err != nil {
		return err
	}

	return nil
}
