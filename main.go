package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"stocknotes/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Could not find .env file!")
		return
	}

	router := chi.NewMux()

	router.Handle("/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr:", listenAddr)

	http.ListenAndServe(listenAddr, router)
}
