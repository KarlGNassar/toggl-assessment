package service

import (
	"context"

	"github.com/KarlGNassar/toggl-assessment/internal/api"
	"github.com/KarlGNassar/toggl-assessment/internal/model"
	"github.com/KarlGNassar/toggl-assessment/internal/store"
	"github.com/KarlGNassar/toggl-assessment/internal/utils"
	"github.com/google/uuid"
)

type DeckService struct {
	store store.DeckStore
}

func NewDeckService(store store.DeckStore) *DeckService {
	return &DeckService{store: store}
}

func (s *DeckService) CreateDeck(ctx context.Context, shuffle bool, cards string) (model.Deck, error) {
	generatedCards := utils.GenerateCards(cards)
	deck := model.Deck{
		Id:        uuid.New(),
		Shuffled:  shuffle,
		Remaining: len(generatedCards),
		Cards:     generatedCards,
	}

	if shuffle {
		utils.ShuffleCards(deck.Cards)
	} else {
		utils.SortDeckCards(deck.Cards)
	}

	err := s.store.CreateDeck(ctx, deck)
	if err != nil {
		return model.Deck{}, err
	}

	return model.Deck{Id: deck.Id, Shuffled: deck.Shuffled, Remaining: deck.Remaining}, nil
}

func (s *DeckService) OpenDeck(ctx context.Context, id uuid.UUID, shuffle bool) (model.Deck, error) {
	deck, err := s.store.GetDeck(ctx, id)
	deck.Shuffled = shuffle
	if err != nil {
		return model.Deck{}, &api.ErrDeckNotFound{Id: id.String()}
	}

	if shuffle {
		utils.ShuffleCards(deck.Cards)
	} // else list the cards as they were created in the database
	return deck, nil
}

func (s *DeckService) DrawCards(ctx context.Context, deckId uuid.UUID, count int) ([]model.Card, error) {
	deck, err := s.store.GetDeck(ctx, deckId)
	if err != nil {
		return nil, &api.ErrDeckNotFound{Id: deckId.String()}
	}

	if count > deck.Remaining {
		return nil, &api.ErrInsufficientCards{Requested: count, Available: deck.Remaining}
	}

	drawnCards := deck.Cards[:count]
	deck.Cards = deck.Cards[count:]
	deck.Remaining -= count

	err = s.store.UpdateDeck(ctx, deckId, deck)
	if err != nil {
		return nil, err
	}
	return drawnCards, nil
}
