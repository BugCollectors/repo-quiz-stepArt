package db_layer

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Chepushila interface {
	SaveMessageFromChepush(ctx context.Context, id, title, message string) error
	GetChepushMessageByID(ctx context.Context, id string) error
	DeleteChepushMessageByID(ctx context.Context, id string) error
}

type ChepushilaRepo struct {
	DB *pgx.Conn
}

func NewChepushilaRepo(connection *pgx.Conn) *ChepushilaRepo {
	return &ChepushilaRepo{
		DB: connection,
	}
}

func (ch *ChepushilaRepo) SaveMessageFromChepush(ctx context.Context, id, title, message string) error {
	var index int
	query := "INSERT INTO message_from_chepush (id, title, message) VALUES ($1, $2, $3) RETURNING id"
	result, err := ch.DB.Query(ctx, query, id, title, message)
	if err != nil {
		return err
	}

	defer result.Close()
	if result.Next() {
		err = result.Scan(&index)
	}
	return nil
}

type ChepushMessage struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}

func (ch *ChepushilaRepo) GetChepushMessageByID(ctx context.Context, id string) (*ChepushMessage, error) {
	message := ChepushMessage{}
	query := `SELECT id, title, message FROM "message_from_chepush" WHERE id = $1`
	result, err := ch.DB.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer result.Close()
	if result.Next() {
		err = result.Scan(&message.ID, &message.Title, &message.Message)
	}
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (ch *ChepushilaRepo) DeleteChepushMessageByID(ctx context.Context, id string) error {
	var index int
	query := "DELETE INTO message_from_chepush WHERE (id) VALUES ($1, $2, $3) RETURNING id"
	result, err := ch.DB.Query(ctx, query, id)
	if err != nil {
		return err
	}

	defer result.Close()
	if result.Next() {
		err = result.Scan(&index)
	}
	return nil
}
