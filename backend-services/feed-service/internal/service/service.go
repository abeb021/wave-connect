package service

import (
	"feed-service/internal/repository"
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

func (s *Service) CreatePublication(ctx context.Context, pubReq repository.PublicationRequest) (*repository.Publication, error) {
	return s.Repo.CreatePublication(ctx, pubReq)
}

func (s *Service) GetFeed(ctx context.Context) ([]repository.Publication, error){
	return s.Repo.GetFeed(ctx)
}

func (s *Service) GetPublicationsByUser(ctx context.Context, userId string) ([]repository.Publication, error){
	return s.Repo.GetPublicationsByUser(ctx, userId)
}


func (s *Service) GetPublication(ctx context.Context, id string) (*repository.Publication, error) {
	return s.Repo.GetPublication(ctx, id)
}

func (s *Service) UpdatePublication(ctx context.Context, id string, text string) error {
	return s.Repo.UpdatePublication(ctx, id, text)

}

func (s *Service) DeletePublication(ctx context.Context, id string) error {
	return s.Repo.DeletePublication(ctx, id)

}
