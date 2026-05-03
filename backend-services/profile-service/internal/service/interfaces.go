package service

import (
	"context"
	"profile-service/internal/domain"
)

type RepositoryInterface interface {
	CreateProfile(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error)
	GetProfile(ctx context.Context, id string) (*domain.Profile, error)
	GetProfileByUsername(ctx context.Context, username string) (*domain.Profile, error)
	UpdateProfile(ctx context.Context, prof *domain.Profile) error
	DeleteProfile(ctx context.Context, id string) error
	UpdateAvatar(ctx context.Context, id string, data []byte) (string, string, error)
	GetAvatar(ctx context.Context, id string) ([]byte, error)
}

type ProducerInterface interface {
	Send(topic, key string, value []byte) error
}