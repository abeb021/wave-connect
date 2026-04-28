package service

import (
	"context"
	"encoding/json"
	"log"
	"profile-service/internal/domain"
	"profile-service/internal/kafka"
	"profile-service/internal/repository"
)

type Service struct {
	Repo     *repository.Repository
	producer *kafka.Producer
}

func NewService(repo *repository.Repository, producer *kafka.Producer) *Service {
	return &Service{
		Repo:     repo,
		producer: producer,
	}
}

func (s *Service) CreateProfile(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error) {
	prof, err := s.Repo.CreateProfile(ctx, profReq, id)
	if err != nil {
		return nil, err
	}
	s.sendProfileUpdatedEvent(prof.ID, prof.Username, prof.Bio, prof.Avatar)
	return prof, nil
}

func (s *Service) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	return s.Repo.GetProfile(ctx, id)
}

func (s *Service) GetProfileByUsername(ctx context.Context, username string) (*domain.Profile, error) {
	return s.Repo.GetProfileByUsername(ctx, username)
}

func (s *Service) UpdateProfile(ctx context.Context, prof *domain.Profile) error {
	err := s.Repo.UpdateProfile(ctx, prof)
	if err != nil {
		return err
	}
	s.sendProfileUpdatedEvent(prof.ID, prof.Username, prof.Bio, prof.Avatar)
	return nil
}

func (s *Service) DeleteProfile(ctx context.Context, id string) error {
	return s.Repo.DeleteProfile(ctx, id)
}

func (s *Service) UpdateAvatar(ctx context.Context, id string, data []byte) error {
	username, bio, err := s.Repo.UpdateAvatar(ctx, id, data)
	if err != nil {
		return err
	}
	s.sendProfileUpdatedEvent(id, username, bio, data)
	return nil
}

func (s *Service) GetAvatar(ctx context.Context, id string) ([]byte, error) {
	return s.Repo.GetAvatar(ctx, id)
}

// kafka events
func (s *Service) sendProfileUpdatedEvent(userID, username, bio string, avatar []byte) {
	event := map[string]interface{}{
		"user_id":  userID,
		"username": username,
		"bio":      bio,
		"avatar":   avatar,
	}
	value, _ := json.Marshal(event)

	err := s.producer.Send("profile-updates", userID, value)
	if err != nil {
		log.Printf("kafka send error:%v\n", err)
	}
}
