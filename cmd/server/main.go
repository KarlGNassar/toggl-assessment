package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KarlGNassar/toggl-assessment/internal/api/handler"
	"github.com/KarlGNassar/toggl-assessment/internal/service"
	"github.com/KarlGNassar/toggl-assessment/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := chi.NewRouter()

	logFile, err := os.OpenFile("../../logs/deck-api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file: ", err)
	}
	log.SetOutput(logFile)
	defer logFile.Close()

	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.Default(), NoColor: true}))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://karlnassar01:1OzfwiLlYzTDu4cZ@cluster0.u0ods0b.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		log.Fatalf("Failed to create MongoDb client: %v", err)
	}

	deckStore := store.NewMongoDeckStore(client, "toggl-assessment", "card-game")
	deckService := service.NewDeckService(deckStore)
	deckHandler := handler.NewDeckHandler(deckService)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Post("/deck", deckHandler.CreateDeck)
	r.Get("/deck/{deckId}", deckHandler.OpenDeck)
	r.Put("/deck/{deckId}/draw/{count}", deckHandler.DrawCards)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
