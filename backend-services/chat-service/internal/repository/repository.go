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
func (ps *Repository) CreateMessage(ctx context.Context, msgReq *MessageRequest) (*Message, error) {
	msg := Message{
		ID: uuid.New().String(),
		Text: msgReq.Text,
		Receiver: msgReq.Receiver,
		Sender: msgReq.Sender,
	}
	row := ps.DB.QueryRow(
		ctx,
		`INSERT INTO messages (id, text, sender, receiver)
	 	 VALUES ($1, $2, $3, $4)
	 	 RETURNING time_sent`,
		msg.ID, msg.Text, msg.Sender, msg.Receiver,
	)

	if err := row.Scan(&msg.TimeSent); err != nil {
		return nil, err
	}

	return &msg, nil
}

func (ps *Repository) GetMessage(ctx context.Context, id string) (Message, error) {

	var msg Message

	row := ps.DB.QueryRow(
		ctx,
		`SELECT id, text, sender, receiver, time_sent
		 FROM messages 
		 WHERE id = $1`,
		id)

	err := row.Scan(&msg.ID, &msg.Text, &msg.Sender, &msg.Receiver, &msg.TimeSent)
	if err != nil {
		return Message{}, err
	}

	return msg, nil
}

func (ps *Repository) UpdateMessage(ctx context.Context, id string, text string) error {

	ct, err := ps.DB.Exec(
		ctx,
		`UPDATE messages
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

func (ps *Repository) DeleteMessage(ctx context.Context, id string) error {
	ct, err := ps.DB.Exec(
		ctx,
		`DELETE FROM messages 
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
