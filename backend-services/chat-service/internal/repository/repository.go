package repository

import (
	"chat-service/internal/config"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error){
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
)
	db, err := pgxpool.New(ctx, connStr)
	if err != nil{
		return nil, err
	}
	return db, nil
}

type Message struct {
	ID       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	TimeSent time.Time `json:"timeSent"`
}

type Repository struct{
	DB *pgxpool.Pool
}

func NewRepository (db *pgxpool.Pool) *Repository{
	return &Repository{
		DB: db,
	}
}
func (ps *Repository)CreateMessage(ctx context.Context, msg Message) (Message, error) {
	msg.ID = uuid.New()
	row := ps.DB.QueryRow(
		ctx, 
		`INSERT INTO messages (id, text, sender, receiver)
	 	 VALUES ($1, $2, $3, $4)
	 	 RETURNING time_sent`, 
		msg.ID, msg.Text, msg.Sender, msg.Receiver,
	)
	
	if err := row.Scan(&msg.TimeSent); err != nil{
		return Message{}, err
	}

	return msg, nil
}

func (ps *Repository)GetMessage(ctx context.Context, id uuid.UUID) (Message, error) {

	var msg Message
	
	row := ps.DB.QueryRow(
		ctx,
		`SELECT id, text, sender, receiver, time_sent
		 FROM messages 
		 WHERE id = $1`, 
		id )
	
	err := row.Scan(&msg.ID, &msg.Text , &msg.Sender, &msg.Receiver, &msg.TimeSent )
	if err != nil{
		if err == pgx.ErrNoRows{
			return Message{}, errors.New("ID not found")
		}

		return Message{}, err
	}

	return msg, nil
}

func (ps *Repository)UpdateMessage(ctx context.Context, id uuid.UUID, text string) (error) {

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

func (ps *Repository)DeleteMessage(ctx context.Context, id uuid.UUID) (error) {
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
