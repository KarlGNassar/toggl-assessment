package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/KarlGNassar/toggl-assessment/internal/api"
	"github.com/KarlGNassar/toggl-assessment/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type DeckHandler struct {
	service *service.DeckService
}

func NewDeckHandler(service *service.DeckService) *DeckHandler {
	return &DeckHandler{service: service}
}

func (h *DeckHandler) CreateDeck(w http.ResponseWriter, r *http.Request) {
	shuffle := r.URL.Query().Get("shuffle") == "true"
	cards := r.URL.Query().Get("cards")

	deck, err := h.service.CreateDeck(r.Context(), shuffle, cards)

	if err != nil {
		log.Printf("Something went wrong while creating the deck: %s", cards)
		api.InternalServerError(w, err)
		return
	}

	api.Created(w, deck)
}

func (h *DeckHandler) OpenDeck(w http.ResponseWriter, r *http.Request) {
	deckId := chi.URLParam(r, "deckId")
	shuffle := r.URL.Query().Get("shuffle") == "true"
	deckIdUUID, err := uuid.Parse(deckId)
	if err != nil {
		log.Printf("Failed to parse UUID: %v", err)
		api.BadRequest(w, api.ParsingFailed, err)
		return
	}

	deck, err := h.service.OpenDeck(r.Context(), deckIdUUID, shuffle)
	if err != nil {
		switch err.(type) {
		case *api.ErrDeckNotFound:
			api.NotFound(w)
			return
		default:
			api.InternalServerError(w, err)
			return
		}
	}

	api.RawOk(w, deck)
}

func (h *DeckHandler) DrawCards(w http.ResponseWriter, r *http.Request) {
	countParam := chi.URLParam(r, "count")
	deckId := chi.URLParam(r, "deckId")
	deckIdUUID, err := uuid.Parse(deckId)
	if err != nil {
		log.Printf("Failed to parse UUID: %v", err)
		api.BadRequest(w, api.ParsingFailed, err)
		return
	}

	count, err := strconv.Atoi(countParam)
	if err != nil || count <= 0 {
		log.Printf("Invalid Count %s", countParam)
		api.BadRequest(w, api.InvalidParam, err)
		return
	}

	cards, err := h.service.DrawCards(r.Context(), deckIdUUID, count)
	if err != nil {
		switch e := err.(type) {
		case *api.ErrDeckNotFound:
			log.Printf("couldn't find deck with id %s", deckId)
			api.NotFound(w)
			return
		case *api.ErrInsufficientCards:
			api.BadRequest(w, api.InsufficientCards, e)
			return
		default:
			api.InternalServerError(w, err)
			return
		}
	}

	api.Ok(w, cards, "cards")
}
