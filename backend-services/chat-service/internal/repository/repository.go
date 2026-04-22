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
		ID:       uuid.New().String(),
		Text:     msgReq.Text,
		Receiver: msgReq.Receiver,
		Sender:   msgReq.Sender,
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

func (ps *Repository) GetConversation(ctx context.Context, senderID string) ([]Message, error) {
	rows, err := ps.DB.Query(
		ctx,
		`SELECT id, text, sender, receiver, time_sent
		 FROM messages
		 WHERE sender = $1  OR receiver = $1
		 ORDER BY time_sent ASC;`,
		senderID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []Message

	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.Text, &msg.Sender, &msg.Receiver, &msg.TimeSent)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return msgs, nil
}

func (ps *Repository) GetConversationWithPeer(ctx context.Context, senderID, receiverID string) ([]Message, error) {
	rows, err := ps.DB.Query(
		ctx,
		`SELECT id, text, sender, receiver, time_sent
		 FROM messages
		 WHERE (sender = $1 AND receiver = $2)
   		 OR (sender = $2 AND receiver = $1)
		 ORDER BY time_sent ASC;`,
		senderID, receiverID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []Message

	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.Text, &msg.Sender, &msg.Receiver, &msg.TimeSent)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return msgs, nil
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

func (ps *Repository) UpdateMessage(ctx context.Context, id, text, senderID string) error {

	ct, err := ps.DB.Exec(
		ctx,
		`UPDATE messages
		 SET text = $1
		 WHERE id = $2 AND sender = $3`,
		text, id, senderID,
	)

	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("message not found")
	}

	return nil
}

func (ps *Repository) DeleteMessage(ctx context.Context, id, senderID string) error {
	ct, err := ps.DB.Exec(
		ctx,
		`DELETE FROM messages 
		 WHERE id=$1 AND sender = $2`,
		id, senderID,
	)

	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("message not found")
	}

	return nil
}
