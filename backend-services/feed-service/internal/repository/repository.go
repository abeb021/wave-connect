package repository

import (
	"context"
	"errors"
	"feed-service/internal/domain"

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
func (ps *Repository) CreatePublication(ctx context.Context, pubReq *domain.PublicationRequest) (*domain.Publication, error) {
	pub := domain.Publication{
		ID:     uuid.New().String(),
		Text:   pubReq.Text,
		UserID: pubReq.UserID,
	}
	row := ps.DB.QueryRow(
		ctx,
		`INSERT INTO publications (id, text, user_id)
	 	 VALUES ($1, $2, $3)
	 	 RETURNING time_created`,
		pub.ID, pub.Text, pub.UserID,
	)

	if err := row.Scan(&pub.TimeCreated); err != nil {
		return nil, err
	}

	return &pub, nil
}

func (ps *Repository) GetFeed(ctx context.Context) ([]domain.Publication, error) {
	rows, err := ps.DB.Query(
		ctx,
		`SELECT id, text, user_id, time_created
		FROM publications
		ORDER BY time_created DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pubs []domain.Publication
	for rows.Next() {
		var pub domain.Publication
		err := rows.Scan(&pub.ID, &pub.Text, &pub.UserID, &pub.TimeCreated)
		if err != nil {
			return nil, err
		}
		pubs = append(pubs, pub)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pubs, nil
}

func (ps *Repository) GetPublicationsByUser(ctx context.Context, userID string) ([]domain.Publication, error) {
	rows, err := ps.DB.Query(
		ctx,
		`SELECT id, text, user_id, time_created
		FROM publications
		WHERE user_id = $1
		ORDER BY time_created DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pubs []domain.Publication
	for rows.Next() {
		var pub domain.Publication
		err := rows.Scan(&pub.ID, &pub.Text, &pub.UserID, &pub.TimeCreated)
		if err != nil {
			return nil, err
		}
		pubs = append(pubs, pub)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pubs, nil
}

func (ps *Repository) GetPublication(ctx context.Context, id string) (*domain.Publication, error) {
	var pub domain.Publication

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

func (ps *Repository) UpdatePublication(ctx context.Context, id, text, userID string) error {

	ct, err := ps.DB.Exec(
		ctx,
		`UPDATE publications
		 SET text = $1
		 WHERE id = $2 AND user_id = $3`,
		text, id, userID,
	)

	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("ID not found")
	}

	return nil
}

func (ps *Repository) DeletePublication(ctx context.Context, id, userID string) error {
	ct, err := ps.DB.Exec(
		ctx,
		`DELETE FROM publications 
		 WHERE id=$1 AND user_id = $2`,
		id, userID,
	)

	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("ID not found")
	}

	return nil
}

//COMMENTS

func (ps *Repository) CreateComment(ctx context.Context, commentReq *domain.CommentRequest) (*domain.Comment, error) {
	comment := domain.Comment{
		ID:     uuid.New().String(),
		PubID:  commentReq.PubID,
		Text:   commentReq.Text,
		UserID: commentReq.UserID,
	}
	row := ps.DB.QueryRow(
		ctx,
		`INSERT INTO comments (id, pub_id, text, user_id)
	 	 VALUES ($1, $2, $3, $4)
	 	 RETURNING time_created`,
		comment.ID, comment.PubID, comment.Text, comment.UserID,
	)

	if err := row.Scan(&comment.TimeCreated); err != nil {
		return nil, err
	}

	return &comment, nil
}

func (ps *Repository) GetCommentsByPublication(ctx context.Context, pubID string) ([]domain.Comment, error) {
	rows, err := ps.DB.Query(
		ctx,
		`SELECT id, pub_id, text, user_id, time_created
		 FROM comments
		 WHERE pub_id = $1
	 	 ORDER BY time_created DESC`,
		 pubID,
	)

	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var comments []domain.Comment
	for rows.Next(){
		var comment domain.Comment
		err := rows.Scan(&comment.ID, &comment.PubID, &comment.Text, &comment.UserID, &comment.TimeCreated)
		if err != nil{
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (ps *Repository) DeleteComment(ctx context.Context, id, userID string) error {
	ct, err := ps.DB.Exec(
		ctx,
		`DELETE FROM comments 
		 WHERE id=$1 AND user_id = $2`,
		id, userID,
	)

	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("ID not found")
	}

	return nil
}