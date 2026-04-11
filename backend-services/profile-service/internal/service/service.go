package service

import (
	"profile-service/internal/repository"
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

func (s *Service) CreateProfile(ctx context.Context, profReq *repository.CreateProfileRequest) (*repository.Profile, error) {
	return s.Repo.CreateProfile(ctx, profReq)
}

func (s *Service) GetProfile(ctx context.Context, id string) (*repository.Profile, error) {
	return s.Repo.GetProfile(ctx, id)
}

func (s *Service) UpdatePublication(ctx context.Context, id string, text string) error {
	return s.Repo.UpdatePublication(ctx, id, text)
}

func (s *Service) DeletePublication(ctx context.Context, id string) error {
	return s.Repo.DeletePublication(ctx, id)

}
