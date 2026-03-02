package main

import (
	"log"
	"net/http"
	"github.com/surajmicky/distributed-rate-limiter/internal/server"
	"github.com/surajmicky/distributed-rate-limiter/internal/redis"
)

func main() {

	// Initialize Redis
	rdb := redis.NewRedisClient()

	if err := redis.Ping(rdb); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	log.Println("Redis connected")

	// Setup HTTP server
	handler := server.RegisterRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Println("Server running on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}