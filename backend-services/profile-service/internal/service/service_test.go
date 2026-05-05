package service

import (
	"context"
	"profile-service/internal/domain"
	"testing"
)

func TestCreateProfile_Success(t *testing.T) {
	repo := &mockRepo{
		CreateProfileMock: func(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error) {
			return &domain.Profile{ID: id, Username: profReq.Username}, nil
		},
	}	
	
	sent := false
	producer := &mockProducer{
		SendMock: func(topic, key string, value []byte) error {
			sent = true
			return nil
		},
	}

	svc := NewService(repo, producer)

	prof, err := svc.CreateProfile(context.Background(), &domain.CreateProfileRequest{Username: "testuser"}, "id123")
	if err != nil{
		t.Fatalf("unexpected error: %v", err)
	}

	if prof.Username != "testuser" {
		t.Errorf("expected testuser, got: %s", prof.Username)
	}

	if !sent{
		t.Error("expected kafka event to be called")
	}
}