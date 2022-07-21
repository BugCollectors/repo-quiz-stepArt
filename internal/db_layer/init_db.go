package db_layer

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"study/internal/cfg"
)

func CreateDB(ctx context.Context, config *cfg.Config) (*pgx.Conn, error) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.PGHost,
		config.PGPort,
		config.PGUsername,
		config.PGPassword,
		config.PGDBName,
		config.PGSSLMode,
	)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
