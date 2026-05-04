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
	DeleteProfileMock        func(ctxMock context.Context, id string) error
	UpdateAvatarMock         func(ctx context.Context, id string, data []byte) (string, string, error)
	GetAvatarMock            func(ctx context.Context, id string) ([]byte, error)
}



func (m *mockRepo) CreateProfile(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error) {
	return m.CreateProfileMock(ctx, profReq, id)
}
func (m *mockRepo) GetProfile(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error) {
	return m.GetProfileMock(ctx, profReq, id)
}
func (m *mockRepo) CreateProfile(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error) {
	return m.CreateProfileMock(ctx, profReq, id)
}
func (m *mockRepo) CreateProfile(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error) {
	return m.CreateProfileMock(ctx, profReq, id)
}
func (m *mockRepo) CreateProfile(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error) {
	return m.CreateProfileMock(ctx, profReq, id)
}