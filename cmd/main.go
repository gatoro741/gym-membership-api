package main

import (
	"GymMembership-api/internal/config"
	"GymMembership-api/internal/storage"
	"context"
	"log"
)

func main() {
	cfg, ctx := config.Load(), context.Background()
	pool := storage.NewDb(ctx, cfg)
	defer pool.Close()

	log.Println("Connected to Database!")
}
