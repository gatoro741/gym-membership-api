package main

import (
	"GymMembership-api/internal/config"
	"GymMembership-api/internal/handlers"
	"GymMembership-api/internal/router"
	"GymMembership-api/internal/service"
	"GymMembership-api/internal/storage"
	"GymMembership-api/internal/worker"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		log.Println("Shutting down...")
		cancel()
	}()

	cfg := config.Load()
	pool := storage.NewDb(ctx, cfg)
	defer pool.Close()

	strg := storage.New(pool)
	svc := service.New(strg)
	h := handlers.New(svc)
	w := worker.New(strg)
	w.Start(ctx)

	r := router.NewRouter(h)

	log.Println("Server starting on 8080")
	http.ListenAndServe(":8080", r)
}
