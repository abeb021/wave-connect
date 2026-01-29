package database

import (
	"context"
	"fmt"
	"chat-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error){
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
)
	db, err := pgxpool.New(ctx, connStr)
	if err != nil{
		return nil, err
	}
	return db, nil
}