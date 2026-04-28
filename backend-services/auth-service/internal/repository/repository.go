package repository

import (
	"auth-service/internal/domain"
	"context"
	"strings"

	_ "github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Repository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		DB: db,
	}
}

func (ps *Repository) Register(ctx context.Context, usr *domain.UserDB) (*domain.UserResponse, error) {
	row := ps.DB.QueryRow(
		ctx,
		`INSERT INTO users (id, email, password_hash)
	 	 VALUES ($1, $2, $3)
		 RETURNING time_created`,
		usr.ID, usr.Email, usr.PasswordHASH,
	)

	if err := row.Scan(&usr.TimeCreated); err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, domain.ErrUserTaken
		}
		return nil, err
	}
	usrResponse := domain.UserResponse{
		ID:        usr.ID,
		Email:     usr.Email,
		CreatedAt: usr.TimeCreated,
	}
	return &usrResponse, nil
}

func (ps *Repository) Login(ctx context.Context, identifier string) (*domain.UserDB, error) {

	if strings.TrimSpace(identifier) == "" {
		return nil, domain.ErrUserNotFound
	}

	var usrDB domain.UserDB
	row := ps.DB.QueryRow(
		ctx,
		`SELECT id, email, password_hash
         FROM users 
         WHERE email = $1
         LIMIT 1`,
		identifier,
	)

	err := row.Scan(&usrDB.ID, &usrDB.Email, &usrDB.PasswordHASH)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &usrDB, nil
}

func (ps *Repository) GetUserById(ctx context.Context, id string) (*domain.UserResponse, error) {
	usr := &domain.UserResponse{}

	row := ps.DB.QueryRow(
		ctx,
		`SELECT id, email, created_at
		 FROM users 
		 WHERE id = $1`,
		id)

	err := row.Scan(&usr.ID, &usr.Email, &usr.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return usr, nil
}

func (ps *Repository) DeleteUser(ctx context.Context, id string) error {
	ct, err := ps.DB.Exec(
		ctx,
		`DELETE FROM users 
		 WHERE id=$1`,
		id,
	)

	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}
