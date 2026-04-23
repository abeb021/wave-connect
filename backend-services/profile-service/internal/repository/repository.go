package repository

import (
	"context"
	"errors"
	"strings"

	"profile-service/internal/domain"

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
func (ps *Repository) CreateProfile(ctx context.Context, profReq *domain.CreateProfileRequest, id string) (*domain.Profile, error) {
	prof := domain.Profile{
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
			return nil, domain.ErrUsernameTaken
		}
		return nil, err
	}

	return &prof, nil
}

func (ps *Repository) GetProfile(ctx context.Context, user_id string) (*domain.Profile, error) {

	var prof domain.Profile

	row := ps.DB.QueryRow(
		ctx,
		`SELECT id, username, bio, time_created
		 FROM profiles 
		 WHERE id = $1`,
		user_id)

	err := row.Scan(&prof.ID, &prof.Username, &prof.Bio, &prof.TimeCreated)
	if err != nil {
		return nil, err
	}

	return &prof, nil
}

func (ps *Repository) GetProfileByUsername(ctx context.Context, username string) (*domain.Profile, error) {
	prof := &domain.Profile{}

	row := ps.DB.QueryRow(
		ctx,
		`SELECT id, username, bio, time_created
		 FROM profiles 
		 WHERE username = $1`,
		username)

	err := row.Scan(&prof.ID, &prof.Username, &prof.Bio, &prof.TimeCreated)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrProfileNotFound
		}
		return nil, err
	}

	return prof, nil
}

func (ps *Repository) UpdateProfile(ctx context.Context, prof *domain.Profile) error {
	ct, err := ps.DB.Exec(
		ctx,
		`UPDATE profiles
		 SET username = $1,
		     bio = $2
		 WHERE id = $3
		 `,
		prof.Username, prof.Bio, prof.ID,
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

func (ps *Repository) UpdateAvatar(ctx context.Context, userID string, data []byte) error{
	ct, err := ps.DB.Exec(ctx,
		`UPDATE profiles
		 SET avatar = $1
		 WHERE id = $2`,
		data, userID,
	)

	
	if err != nil{
		return err
	}
	
	if ct.RowsAffected() == 0{
		return domain.ErrProfileNotFound
	}

	return nil
}


func (ps *Repository) GetAvatar(ctx context.Context, userID string) ([]byte, error) {
	row := ps.DB.QueryRow(
		ctx,
		`SELECT avatar
		 FROM profiles 
		 WHERE id = $1`,
		userID,
	)
	var imgBytes []byte
	err := row.Scan(&imgBytes)
	if err != nil {
		return nil, err
	}

	return imgBytes, nil
}
