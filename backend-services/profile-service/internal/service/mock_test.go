package service

import (
	"context"
	"profile-service/internal/domain"
)

type mockRepo struct {
	CreateProfileMock        func(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error)
	GetProfileMock           func(ctx context.Context, id string) (*domain.Profile, error)
	GetProfileByUsernameMock func(ctx context.Context, username string) (*domain.Profile, error)
	UpdateProfileMock        func(ctx context.Context, prof *domain.Profile) error
	DeleteProfileMock        func(ctx context.Context, id string) error
	UpdateAvatarMock         func(ctx context.Context, id string, data []byte) (string, string, error)
	GetAvatarMock            func(ctx context.Context, id string) ([]byte, error)
}

func (m *mockRepo) CreateProfile(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error) {
	return m.CreateProfileMock(ctx, profReq, id)
}
func (m *mockRepo) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	return m.GetProfileMock(ctx, id)
}
func (m *mockRepo) GetProfileByUsername(ctx context.Context, username string) (*domain.Profile, error) {
	return m.GetProfileByUsernameMock(ctx, username)
}
func (m *mockRepo) UpdateProfile(ctx context.Context, prof *domain.Profile) error {
	return m.UpdateProfileMock(ctx, prof)
}
func (m *mockRepo) DeleteProfile(ctx context.Context, id string) error{
	return m.DeleteProfileMock(ctx, id)
}
func (m *mockRepo) UpdateAvatar(ctx context.Context, id string, data []byte) (string, string, error) {
	return m.UpdateAvatarMock(ctx, id, data)
}
func (m *mockRepo) GetAvatar(ctx context.Context, id string) ([]byte, error) {
	return m.GetAvatarMock(ctx, id)
}

type mockProducer struct {
	SendMock func(topic, key string, value []byte) error
}

func (m *mockProducer) Send(topic, key string, value []byte) error{
	return m.SendMock(topic, key, value)
}