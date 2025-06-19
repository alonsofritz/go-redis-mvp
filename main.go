package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Nome do serviço no docker-compose
		Password: "",
		DB:       0, // Default DB
	})

	// Testar conexão com Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")
}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	counter, err := rdb.Incr(ctx, "counter").Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Counter value: %d\n", counter)
	log.Printf("Counter incremented to: %d", counter)
}

func main() {
	initRedis()

	http.HandleFunc("/increment", incrementHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
