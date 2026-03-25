package services

import (
	"context"
	"easy-ride-api/internal/domain"
	entity "easy-ride-api/internal/models"
	"easy-ride-api/internal/respositories"
	"easy-ride-api/pkg/utils"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateNewUser(ctx context.Context, fullName, email, password, confirmPassword string) (*entity.User, error)
	CreateUserSession(ctx context.Context, email, password string) (*entity.Session, error)
}

type userService struct {
	userRepo respositories.UserRepository
}

func NewUserService(userRepo respositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateNewUser(ctx context.Context, fullName, email, password, confirmPassword string) (*entity.User, error) {
	if password != confirmPassword {
		return nil, errors.New("passwords do not match")
	}

	existing, _ := s.userRepo.GetUserByEmail(ctx, email)
	if existing != nil {
		return nil, errors.New("email already in use")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.CreateUser(ctx, &entity.User{
		FullName: fullName,
		Email:    email,
		Password: string(hashedPass),
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUserSession(c context.Context, email, password string) (*entity.Session, error) {
	user, err := s.userRepo.GetUserByEmail(c, email)

	if err != nil {
		return nil, err
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if passErr != nil {
		return nil, passErr
	}

	//Build up session struct
	token, err := utils.GenerateSecureToken(24)

	if err != nil {
		return nil, errors.New("Failed to generate token")
	}

	now := time.Now()

	expiry := now.Add(time.Duration(domain.SESSION_EXPIRY_IN_SECONDS) * time.Second)

	session := &entity.Session{
		UserId:    user.ID,
		Token:     token,
		ExpiresAt: expiry,
		CreatedAt: now,
		UpdatedAt: now,
	}

	session, sessionErr := s.userRepo.CreateSession(c, session)

	if sessionErr != nil {
		return nil, sessionErr
	}

	return session, nil
}
