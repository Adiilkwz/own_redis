package main

import (
	"log"

	"own-redis/internal/config"
	"own-redis/internal/protocol"
	"own-redis/internal/server"
	"own-redis/internal/storage"
)

func main() {
	cfg := config.ParseFlags()

	store := storage.NewStore()
	handler := protocol.NewHandler(store)

	srv := server.NewUDPServer(cfg.Port, handler)
	if err := srv.Start(); err != nil {
		log.Fatalf("Fatal server error: %v\n", err)
	}
}
