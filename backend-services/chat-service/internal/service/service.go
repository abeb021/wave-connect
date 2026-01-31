package service

import (
	"chat-service/internal/repository"
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

func (s *Service) CreateMessage(ctx context.Context, msg repository.Message) (repository.Message, error) {
	return s.Repo.CreateMessage(ctx, msg)
}

func (s *Service) GetMessage(ctx context.Context, id string) (repository.Message, error) {
	return s.Repo.GetMessage(ctx, id)

}

func (s *Service) UpdateMessage(ctx context.Context, id string, text string) error {
	return s.Repo.UpdateMessage(ctx, id, text)

}

func (s *Service) DeleteMessage(ctx context.Context, id string) error {
	return s.Repo.DeleteMessage(ctx, id)

}
