package database

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(context.Context) (*pgxpool.Pool, error){
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"), 
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil{
		return nil, err
	}
	return db, nil
}