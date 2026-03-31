package services

import (
	"context"
	"easy-ride-api/internal/models"
	repositories "easy-ride-api/internal/repositories"
)

type SessionService interface {
	GetSession(c context.Context, token string) (*models.Session, error)
}

type sessionService struct {
	sessionRepo repositories.SessionRepo
}

func NewSessionService(s repositories.SessionRepo) *sessionService {
	return &sessionService{
		sessionRepo: s,
	}
}

func (s *sessionService) GetSession(c context.Context, token string) (*models.Session, error) {
	session, err := s.sessionRepo.GetSessionByToken(c, token)

	if err != nil {
		return nil, err
	}
	return session, nil
}
