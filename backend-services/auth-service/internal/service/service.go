package service

import (
	"auth-service/internal/repository"
	"auth-service/util"
	"context"
)

type Service struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) Register(ctx context.Context, usr repository.User) (string, error) {
	hashedPassword, err := util.HashPassword(usr.Password)
	if err != nil {
		return "", err
	}
	return s.Repo.Register(ctx, usr{Password:hashedPassword})
}

func (s *Service) GetUserById(ctx context.Context, id string) (repository.User, error) {
	return s.Repo.GetUserById(ctx, id)

}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.Repo.DeleteUser(ctx, id)

}
