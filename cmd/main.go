package main

import (
	"GymMembership-api/internal/config"
	"GymMembership-api/internal/service"
	"GymMembership-api/internal/storage"
	"context"
	"log"
)

func main() {
	cfg, ctx := config.Load(), context.Background()
	pool := storage.NewDb(ctx, cfg)
	defer pool.Close()

	strg := storage.New(pool)
	serv := service.New(strg)

	log.Println("Connected to Database!")
}
