package service

import (
	"context"
	"profile-service/internal/repository"
)

type Service struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateProfile(ctx context.Context, profReq *repository.CreateProfileRequest, id string) (*repository.Profile, error) {
	return s.Repo.CreateProfile(ctx, profReq, id)
}

func (s *Service) GetProfile(ctx context.Context, id string) (*repository.Profile, error) {
	return s.Repo.GetProfile(ctx, id)
}

func (s *Service) UpdateProfile(ctx context.Context, prof *repository.Profile) error {
	return s.Repo.UpdateProfile(ctx, prof)
}

func (s *Service) DeleteProfile(ctx context.Context, id string) error {
	return s.Repo.DeleteProfile(ctx, id)
}

func (s *Service) UpdateAvatar(ctx context.Context, id string, data []byte) error {
	return s.Repo.UpdateAvatar(ctx, id, data)
}

func (s *Service) GetAvatar(ctx context.Context, id string) ([]byte, error) {
	return s.Repo.GetAvatar(ctx, id)
}