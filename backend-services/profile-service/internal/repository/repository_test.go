package repository

import ( 
	"context"
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