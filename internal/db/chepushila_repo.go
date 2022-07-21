package db

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Chepushila interface {
	WriteMessageFromChepush(ctx context.Context, title string, message string) error
}

type ChepushilaRepo struct {
	DB *pgx.Conn
}

func NewChepushilaRepo(connection *pgx.Conn) *ChepushilaRepo {
	return &ChepushilaRepo{
		DB: connection,
	}
}

func (ch *ChepushilaRepo) WriteMessageFromChepush(ctx context.Context, title string, message string) error {
	var id int
	query := "INSERT INTO message_from_chepush (title) VALUES ($1) (message) VALUES ($2) RETURNING id"
	result, err := ch.DB.Query(ctx, query, title, message)
	if err != nil {
		return err
	}

	defer result.Close()
	if result.Next() {
		err = result.Scan(&id)
	}
	return nil
}

//id int NOT NULL,
//title text,
//body text,
//PRIMARY KEY(id)
//
//result, err := tx.Query(ctx, query, displayName)
//if err == nil {
//defer result.Close()
//if result.Next() {
//err = result.Scan(&id)
//}
//}
//return nil
