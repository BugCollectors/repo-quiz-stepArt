package db

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

func (ch *ChepushilaRepo) GetChepushMessageByID(ctx context.Context, id, title, message string) error {
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

func (ch *ChepushilaRepo) DeleteChepushMessageByID(ctx context.Context, id, title, message string) error {
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