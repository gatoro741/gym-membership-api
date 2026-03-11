package main

import (
	"GymMembership-api/internal/config"
	"GymMembership-api/internal/handlers"
	"GymMembership-api/internal/router"
	"GymMembership-api/internal/service"
	"GymMembership-api/internal/storage"
	"context"
	"log"
	"net/http"
)

func main() {
	cfg, ctx := config.Load(), context.Background()
	pool := storage.NewDb(ctx, cfg)
	defer pool.Close()

	strg := storage.New(pool)
	svc := service.New(strg)
	h := handlers.New(svc)

	r := router.NewRouter(h)

	log.Println("Server starting on 8080")
	http.ListenAndServe(":8080", r)
}
