package utils

import (
	"reflect"
	"testing"

	"github.com/KarlGNassar/toggl-assessment/internal/constants"
	"github.com/KarlGNassar/toggl-assessment/internal/model"
	"github.com/google/uuid"
)

func TestShuffleCards(t *testing.T) {
	cards := []model.Card{
		{Value: "Ace", Suit: "Spades", Code: "AS"},
		{Value: "Two", Suit: "Hearts", Code: "2H"},
		{Value: "King", Suit: "Clubs", Code: "KC"},
		{Value: "Ten", Suit: "Diamonds", Code: "10D"},
	}

	ShuffleCards(cards)

	if len(cards) != 4 {
		t.Errorf("Expected 4 cards, got %d", len(cards))
	}

	expectedCards := map[model.Card]bool{
		{Value: "Ace", Suit: "Spades", Code: "AS"}:    true,
		{Value: "Two", Suit: "Hearts", Code: "2H"}:    true,
		{Value: "King", Suit: "Clubs", Code: "KC"}:    true,
		{Value: "Ten", Suit: "Diamonds", Code: "10D"}: true,
	}

	for _, card := range cards {
		if !expectedCards[card] {
			t.Errorf("Unexpected card found after shuffle: %+v", card)
		}
		delete(expectedCards, card)
	}

	if len(expectedCards) != 0 {
		t.Errorf("Not all original cards found after shuffle")
	}

	empty := []model.Card{}
	ShuffleCards(empty)
	if len(empty) != 0 {
		t.Errorf("Shuffling an empty slice should not change its length")
	}

	single := []model.Card{{Value: "Three", Suit: "Hearts", Code: "3H"}}
	originalSingle := make([]model.Card, len(single))
	copy(originalSingle, single)

	ShuffleCards(single)
	if !reflect.DeepEqual(single, originalSingle) {
		t.Errorf("Shuffling a single-element slice should not change the element")
	}
}

func TestGenerateFullDeck(t *testing.T) {
	fullDeck := generateFullDeck()
	expectedDeckSize := len(constants.CardMap)

	if len(fullDeck) != expectedDeckSize {
		t.Errorf("Expected full deck size of %d, but got %d", expectedDeckSize, len(fullDeck))
	}
}

func TestGenerateCards(t *testing.T) {
	fullDeck := GenerateCards("")
	if len(fullDeck) != len(constants.CardMap) {
		t.Errorf("Expected full deck to be generated, but got %d cards", len(fullDeck))
	}

	testCodes := "AS,KD,QC"
	expectedCardsCount := 3

	specificCards := GenerateCards(testCodes)
	if len(specificCards) != expectedCardsCount {
		t.Errorf("Expected %d specific cards, but got %d", expectedCardsCount, len(specificCards))
	}

	nonExistingCodes := "ZZ,XX"
	expectedCardsCount = 0

	nonExistingCards := GenerateCards(nonExistingCodes)
	if len(nonExistingCards) != expectedCardsCount {
		t.Errorf("Expected %d cards for non-existing codes, but got %d", expectedCardsCount, len(nonExistingCards))
	}
}

var testCards = []model.Card{
	{Value: "ACE", Suit: "SPADES", Code: "AS"},
	{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
	{Value: "QUEEN", Suit: "CLUBS", Code: "QC"},
	{Value: "JACK", Suit: "HEARTS", Code: "JH"},
	{Value: "10", Suit: "SPADES", Code: "10S"},
}

func TestCardSortIndexReflect(t *testing.T) {
	tests := []struct {
		card     model.Card
		expected int
	}{
		{testCards[0], 1},   // ACE of SPADES
		{testCards[1], 113}, // KING of DIAMONDS
		{testCards[2], 212}, // QUEEN of CLUBS
		{testCards[3], 311}, // JACK of HEARTS
		{testCards[4], 10},  // 10 of SPADES
	}

	for _, test := range tests {
		actual := cardSortIndex(test.card)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("cardSortIndex(%v) = %v, want %v", test.card, actual, test.expected)
		}
	}
}

func TestSortDeckCardsReflect(t *testing.T) {
	deck := model.Deck{
		Id:        uuid.New(),
		Shuffled:  false,
		Remaining: len(testCards),
		Cards:     []model.Card{testCards[1], testCards[4], testCards[2], testCards[0], testCards[3]}, // Unsorted
	}

	expectedOrder := []model.Card{testCards[0], testCards[4], testCards[1], testCards[2], testCards[3]} // Sorted

	SortDeckCards(deck.Cards)

	if !reflect.DeepEqual(deck.Cards, expectedOrder) {
		t.Errorf("SortDeckCards() resulted in %v, want %v", deck.Cards, expectedOrder)
	}
}
