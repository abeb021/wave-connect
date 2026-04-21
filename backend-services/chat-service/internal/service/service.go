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

func (s *Service) CreateMessage(ctx context.Context, msg *repository.MessageRequest) (*repository.Message, error) {
	return s.Repo.CreateMessage(ctx, msg)
}

func (s *Service) GetConversation(ctx context.Context, senderID string) ([]repository.Message, error) {
	return s.Repo.GetConversation(ctx, senderID)
}

func (s *Service) GetConversationWithPeer(ctx context.Context, senderID, receiverID string) ([]repository.Message, error) {
	return s.Repo.GetConversationWithPeer(ctx, senderID, receiverID)
}

func (s *Service) GetMessage(ctx context.Context, id string) (repository.Message, error) {
	return s.Repo.GetMessage(ctx, id)
}

func (s *Service) UpdateMessage(ctx context.Context, id string, text string) error {
	return s.Repo.UpdateMessage(ctx, id, text)
}

func (s *Service) DeleteMessage(ctx context.Context, id, senderID string) error {
	return s.Repo.DeleteMessage(ctx, id, senderID)
}
