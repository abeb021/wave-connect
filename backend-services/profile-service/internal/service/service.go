package service

import (
	"context"
	"profile-service/internal/domain"
	"profile-service/internal/kafka"
	"profile-service/internal/repository"
)

type Service struct {
	Repo *repository.Repository
	producer *kafka.Producer
}

func NewService(repo *repository.Repository, producer *kafka.Producer) *Service {
	return &Service{
		Repo: repo,
		producer: producer,
	}
}

func (s *Service) CreateProfile(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error) {
	return s.Repo.CreateProfile(ctx, profReq, id)
}

func (s *Service) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	return s.Repo.GetProfile(ctx, id)
}

func (s *Service) GetProfileByUsername(ctx context.Context, username string) (*domain.Profile, error) {
	return s.Repo.GetProfileByUsername(ctx, username)
}

func (s *Service) UpdateProfile(ctx context.Context, prof *domain.Profile) error {
	err := s.Repo.UpdateProfile(ctx, prof)
	if err != nil{
		return err
	}
	s.producer.Send()

	return nil
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


//kafka events
func (s *Service) sendProfilUpdatedEvent(userID, username string, avatar []byte) {
	
}