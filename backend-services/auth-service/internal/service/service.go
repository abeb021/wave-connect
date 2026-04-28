package service

import (
	"auth-service/internal/repository"
	"auth-service/internal/domain"
	"auth-service/pkg/jwt"
	"auth-service/pkg/util"
	"context"

	"github.com/google/uuid"
)

type Service struct {
	Repo *repository.Repository
	Auth *jwt.AuthService
}

func NewService(repo *repository.Repository, jwt *jwt.AuthService) *Service {
	return &Service{
		Repo: repo,
		Auth: jwt,
	}
}

func (s *Service) Register(ctx context.Context, usrRequest *domain.UserRequest) (*domain.UserResponse, error) {
	hashedPassword, err := util.HashPassword(usrRequest.Password)
	if err != nil {
		return nil, err
	}
	usr := &domain.UserDB{
		ID:           uuid.New().String(),
		Email:        usrRequest.Email,
		PasswordHASH: hashedPassword,
	}

	usrResponse, err := s.Repo.Register(ctx, usr)
	if err != nil {
		return nil, err
	}

	return usrResponse, nil
}

func (s *Service) Login(ctx context.Context, usrRequest *domain.UserRequest) (string, error) {
	usrDB, err := s.Repo.Login(ctx, usrRequest.Email)
	if err != nil {
		return "", err
	}

	err = util.ValidatePassword(usrRequest.Password, usrDB.PasswordHASH)
	if err != nil {
		return "", domain.ErrWrongPassword
	}

	token, err := s.Auth.GenerateToken(usrDB.ID, usrDB.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GetUserById(ctx context.Context, id string) (*domain.UserResponse, error) {
	return s.Repo.GetUserById(ctx, id)
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.Repo.DeleteUser(ctx, id)
}
