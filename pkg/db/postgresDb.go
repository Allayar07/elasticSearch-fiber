package db

import (
	"context"
	"elasticSearch/internal/configs"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/url"
)

func NewPostgres(cfg *configs.Configs) (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=%s", cfg.Postgres.UserName, url.QueryEscape(cfg.Postgres.Password), cfg.Postgres.Host, cfg.Postgres.DbName, cfg.Postgres.Sslmode)
	DbPool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	if err = DbPool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return DbPool, nil
}
