package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
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
func (ps *Repository) CreatePublication(ctx context.Context, pubReq PublicationRequest) (*Publication, error) {
	pub := Publication{
		ID: uuid.New().String(),
		Text: pubReq.Text,
		UserID: pubReq.UserID,
	}
	row := ps.DB.QueryRow(
		ctx,
		`INSERT INTO publications (id, text, user_id)
	 	 VALUES ($1, $2, $3)
	 	 RETURNING time_sent`,
		pub.ID, pub.Text, pub.UserID,
	)

	if err := row.Scan(&pub.TimeCreated); err != nil {
		return nil, err
	}

	return &pub, nil
}

func (ps *Repository) GetPublication(ctx context.Context, id string) (*Publication, error) {
	var pub Publication

	row := ps.DB.QueryRow(
		ctx,
		`SELECT id, text, user_id, time_created
		 FROM publications
		 WHERE id = $1`,
		id)

	err := row.Scan(&pub.ID, &pub.Text, &pub.UserID, &pub.TimeCreated)
	if err != nil {
		return nil, err
	}

	return &pub, nil
}

func (ps *Repository) UpdatePublication(ctx context.Context, id string, text string) error {

	ct, err := ps.DB.Exec(
		ctx,
		`UPDATE publications
		 SET text = $1
		 WHERE id = $2`,
		text, id,
	)

	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("ID not found")
	}

	return nil
}

func (ps *Repository) DeletePublication(ctx context.Context, id string) error {
	ct, err := ps.DB.Exec(
		ctx,
		`DELETE FROM publications 
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
