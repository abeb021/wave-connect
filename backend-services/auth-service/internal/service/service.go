package service

import (
	"auth-service/internal/repository"
	"auth-service/jwt"
	"auth-service/usecases"
	"auth-service/util"
	"context"
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

func (s *Service) Register(ctx context.Context, usrRequest *repository.UserRequest) (*repository.UserResponse, error) {
	hashedPassword, err := util.HashPassword(usrRequest.Password)
	if err != nil {
		return &repository.UserResponse{}, err
	}
	usr := &repository.UserDB{
		ID: uuid.New().String(),
		Username: usrRequest.Username,
		Email: usrRequest.Email,
		PasswordHASH: hashedPassword,
		CreatedAt: time.Now(),
	}

	usrResponse, err := s.Repo.Register(ctx, usr)
	if err != nil{
		return &repository.UserResponse{}, err
	}

	return usrResponse, nil
}

func (s *Service) Login(ctx context.Context, usrRequest *repository.UserRequest) (string, error) {
	usrDB, err := s.Repo.Login(ctx, usrRequest.Username)
	if err != nil{
		return "", err
	}

	err = util.ValidatePassword(usrRequest.Password, usrDB.PasswordHASH)
	if err != nil{
		return "", usecases.ErrWrongPassword
	}
	
	token, err := s.Auth.GenerateToken(usrDB.ID, usrDB.Email) 
	if err != nil {
		return "", err
	}


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