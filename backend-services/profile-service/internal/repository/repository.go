package repository

import (
	"context"
	"errors"
	"strings"

	"profile-service/usecases"
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
func (ps *Repository) CreateProfile(ctx context.Context, profReq *CreateProfileRequest, id string) (*Profile, error) {
	prof := Profile{
		Username: profReq.Username,
		ID:       id,
	}
	row := ps.DB.QueryRow(
		ctx,
		`INSERT INTO profiles (id, username)
	 	 VALUES ($1, $2)
	 	 RETURNING time_created`,
		prof.ID, prof.Username,
	)

	if err := row.Scan(&prof.TimeCreated); err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, usecases.ErrUserTaken
		}
		return nil, err
	}

	return &prof, nil
}

func (ps *Repository) GetProfile(ctx context.Context, user_id string) (*Profile, error) {

	var prof Profile

	row := ps.DB.QueryRow(
		ctx,
		`SELECT id, username, time_created
		 FROM profiles 
		 WHERE id = $1`,
		user_id)

	err := row.Scan(&prof.ID, &prof.Username, &prof.TimeCreated)
	if err != nil {
		return nil, err
	}

	return &prof, nil
}

func (ps *Repository) UpdateProfile(ctx context.Context, prof *Profile) error {
	ct, err := ps.DB.Exec(
		ctx,
		`UPDATE profiles
		 SET username = $1
		 WHERE id = $2
		 `,
		prof.Username, prof.ID,
	)

	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("ID/Username not found")
	}

	return nil
}

func (ps *Repository) DeleteProfile(ctx context.Context, id string) error {
	ct, err := ps.DB.Exec(
		ctx,
		`DELETE FROM profiles 
		 WHERE id=$1`,
		id,
	)

	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("ID not found")
	}

	return nil
}
