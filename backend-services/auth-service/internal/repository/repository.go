package repository

import (
	"auth-service/usecases"
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

func (ps *Repository) Register(ctx context.Context, usr *UserDB) (*UserResponse, error) {
	_, err := ps.DB.Exec(
		ctx,
		`INSERT INTO users (id, username, email, password_hash)
	 	 VALUES ($1, $2, $3, $4)`,
		usr.ID, usr.Username, usr.Email, usr.PasswordHASH,
	)

	if err != nil {
		if strings.Contains(err.Error(), "dublicate") {
			return nil, usecases.ErrUserTaken
		}
		return nil, err
	}
	usrResponse := UserResponse{
		ID: usr.ID,
		Username: usr.Username,
		Email: usr.Email,
		CreatedAt: usr.CreatedAt,
	}
	return &usrResponse, nil
}

func (ps *Repository) Login(ctx context.Context, identifier string) (*UserDB, error) {
    
	if strings.TrimSpace(identifier) == "" {
        return nil, usecases.ErrUserNotFound
    }

	var usrDB UserDB    
    row := ps.DB.QueryRow(
        ctx,
        `SELECT id, username, email, password_hash
         FROM users 
         WHERE username = $1 OR email = $1
         LIMIT 1`,
        identifier,
    )

    err := row.Scan(&usrDB.ID, &usrDB.Username, &usrDB.Email, &usrDB.PasswordHASH)
    
    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, usecases.ErrUserNotFound
        }
        return nil, err
    }
    
    return &usrDB, nil
}

/*
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
*/