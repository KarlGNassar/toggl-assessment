package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KarlGNassar/toggl-assessment/internal/api"
	"github.com/KarlGNassar/toggl-assessment/internal/api/handler"
	"github.com/KarlGNassar/toggl-assessment/internal/service"
	"github.com/KarlGNassar/toggl-assessment/internal/store"
	"github.com/KarlGNassar/toggl-assessment/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := chi.NewRouter()

	err := utils.LoadEnv(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	logFile, err := os.OpenFile("logs/deck-api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file: ", err)
	}
	log.SetOutput(logFile)
	defer logFile.Close()

	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.Default(), NoColor: true}))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatalf("Failed to create MongoDb client: %v", err)
	}

	deckStore := store.NewMongoDeckStore(client, "toggl-assessment", "card-game")
	deckService := service.NewDeckService(deckStore)
	deckHandler := handler.NewDeckHandler(deckService)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/deck", func(r chi.Router) {
		r.Use(api.AuthenticationMiddleware)

		r.Post("/", deckHandler.CreateDeck)
		r.Get("/{deckId}", deckHandler.OpenDeck)
		r.Put("/{deckId}/draw/{count}", deckHandler.DrawCards)
	})

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
