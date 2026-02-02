package service

import (
	"auth-service/internal/repository"
	"auth-service/jwt"
	"auth-service/util"
	"context"
	"log"
	"time"

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

func (s *Service) Register(ctx context.Context, usrRequest repository.UserRequest) (string, error) {
	hashedPassword, err := util.HashPassword(usrRequest.Password)
	if err != nil {
		return "", err
	}
	usr := &repository.UserDB{
		ID: uuid.New().String(),
		Username: usrRequest.Username,
		Email: usrRequest.Email,
		PasswordHASH: hashedPassword,
		CreatedAt: time.Now(),
	}

	_, err = s.Repo.Register(ctx, usr)
	if err != nil{
		return "", err
	}

	token, err := s.Auth.GenerateToken(usr.ID, usr.Email)
	if err != nil{
		return "", err
	}

	log.Println(usr.ID)
	return token, nil
}
/*
func (s *Service) GetUserById(ctx context.Context, id string) (repository.UserResponse, error) {
	return s.Repo.GetUserById(ctx, id)

}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.Repo.DeleteUser(ctx, id)

}
*/