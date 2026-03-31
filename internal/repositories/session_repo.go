package respositories

import (
	"context"
	entity "easy-ride-api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionRepo interface {
	GetSessionByToken(c context.Context, token string) (*entity.Session, error)
}

type sessionRepo struct {
	pool *pgxpool.Pool
}

func NewSessionRepo(p *pgxpool.Pool) *sessionRepo {
	return &sessionRepo{
		pool: p,
	}
}

func (s *sessionRepo) GetSessionByToken(c context.Context, token string) (*entity.Session, error) {

	query := `SELECT id, user_id, token, expires_at, created_at FROM sessions WHERE token = $1 AND expires_at > NOW()`

	session := &entity.Session{}

	err := s.pool.QueryRow(c, query, token).Scan(
		&session.ID,
		&session.UserId,
		&session.Token,
		&session.ExpiresAt,
		&session.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return session, nil
}
