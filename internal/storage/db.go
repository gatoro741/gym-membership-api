package storage

import (
	"GymMembership-api/internal/config"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDb(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatal("Unable to create connection pool: %v\n", err)
	}

	return pool
}
