package utils

import (
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/KarlGNassar/toggl-assessment/internal/constants"
	"github.com/KarlGNassar/toggl-assessment/internal/model"
)

func cardSortIndex(card model.Card) int {
	suitIndex := constants.SuitsOrder[card.Suit]
	valueIndex := constants.ValuesOrder[card.Value]
	return suitIndex*100 + valueIndex // we multiply by 100 for suit priority
}

func SortDeckCards(cards []model.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cardSortIndex(cards[i]) < cardSortIndex(cards[j])
	})
}

func generateFullDeck() []model.Card {
	fullDeck := make([]model.Card, 0, len(constants.CardMap))
	for _, card := range constants.CardMap {
		fullDeck = append(fullDeck, card)
	}
	return fullDeck
}

func GenerateCards(cardsQuery string) []model.Card {
	if cardsQuery == "" {
		return generateFullDeck()
	}

	codes := strings.Split(cardsQuery, ",")
	cards := make([]model.Card, 0, len(codes))
	for _, code := range codes {
		if card, exists := constants.CardMap[code]; exists {
			cards = append(cards, card)
		}
	}
	return cards
}

func ShuffleCards(cards []model.Card) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(cards) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
}
