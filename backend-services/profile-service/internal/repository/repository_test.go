package repository

import (
	"context"
	"profile-service/internal/domain"
	"testing"
)
func TestNewPool(t *testing.T) {
	ctx := context.Background()

	dbURL := "postgres://postgres:password@profile_db:5432/profile_db?sslmode=disable"
	pool, err := NewPool(ctx, dbURL)
	if err != nil{
		t.Fatalf("failed to connect: %v", err)
	}
	defer pool.Close()

	err = pool.Ping(ctx) 
	if err != nil{
		t.Fatalf("failed to ping: %v", err)
	}
}

func TestCreateProfile(t *testing.T) {
	ctx := context.Background()
	dbURL := "postgres://postgres:password@profile_db:5432/profile_db?sslmode=disable"
	pool, err := NewPool(ctx, dbURL)
	if err != nil{
		t.Fatalf("failed to connect: %v", err)
	}
	defer pool.Close()



	repo := NewRepository(pool)

	profReq := &domain.CreateProfileRequest{
		Username: "testuser",
	}
	
	id :=  "00000000-0000-0000-0000-000000000001"

	prof, err := repo.CreateProfile(ctx, profReq, id)
	if err != nil{
		t.Fatalf("CreateProfile failed: %v", err)
	}

    if prof.ID != id {
        t.Errorf("expected ID %s, got %s", id, prof.ID)
    }
    if prof.Username != "testuser" {
        t.Errorf("expected username testuser, got %s", prof.Username)
    }
}

