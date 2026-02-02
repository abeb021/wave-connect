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
func (ps *Repository) Register(ctx context.Context, usr User) (User, error) {
	usr.ID = uuid.New().String()
	row := ps.DB.QueryRow(
		ctx,
		`INSERT INTO users (id, name, email, password)
	 	 VALUES ($1, $2, $3, $4)
	 	 RETURNING time_created`,
		usr.ID, usr.Name, usr.Email, usr.Password,
	)

	if err := row.Scan(&usr.TimeCreated); err != nil {
		return User{}, err
	}

	return usr, nil
}

func (ps *Repository) GetUserById(ctx context.Context, id string) (User, error) {

	var usr User

	row := ps.DB.QueryRow(
		ctx,
		`SELECT id, name, password, email, time_created
		 FROM users 
		 WHERE id = $1`,
		id)

	err := row.Scan(&usr.ID, &usr.Name, &usr.Password, &usr.Email, &usr.TimeCreated)
	if err != nil {
		return User{}, err
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
		return errors.New("ID not found")
	}

	return nil
}
