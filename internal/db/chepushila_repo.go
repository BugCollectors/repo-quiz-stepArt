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

func (*ChepushilaRepo) WriteMessageFromChepush(ctx context.Context, title string, message string) error {
	println("d")
	return nil
}
