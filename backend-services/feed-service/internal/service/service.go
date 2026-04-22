package service

import (
	"context"
	"feed-service/internal/repository"
)

type Service struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreatePublication(ctx context.Context, pubReq *repository.PublicationRequest) (*repository.Publication, error) {
	return s.Repo.CreatePublication(ctx, pubReq)
}

func (s *Service) GetFeed(ctx context.Context) ([]repository.Publication, error) {
	return s.Repo.GetFeed(ctx)
}

func (s *Service) GetPublicationsByUser(ctx context.Context, userId string) ([]repository.Publication, error) {
	return s.Repo.GetPublicationsByUser(ctx, userId)
}

func (s *Service) GetPublication(ctx context.Context, id string) (*repository.Publication, error) {
	return s.Repo.GetPublication(ctx, id)
}

func (s *Service) UpdatePublication(ctx context.Context, id, text, userID string) error {
	return s.Repo.UpdatePublication(ctx, id, text, userID)
}

func (s *Service) DeletePublication(ctx context.Context, id, userID string) error {
	return s.Repo.DeletePublication(ctx, id, userID)
}

func (s *Service) CreateComment(ctx context.Context, commentReq *repository.CommentRequest) (*repository.Comment, error){
	return s.Repo.CreateComment(ctx, commentReq)
}

func (s *Service) GetCommentsByPublication(ctx context.Context, pubID string) ([]repository.Comment, error){
	return s.Repo.GetCommentsByPublication(ctx, pubID)
}
func (s *Service) DeleteComment(ctx context.Context, id, userID string) error{
	return s.Repo.DeleteComment(ctx, id, userID)
}