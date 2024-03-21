package constants

import "github.com/KarlGNassar/toggl-assessment/internal/model"

var CardMap = map[string]model.Card{
	"AS": {Value: "ACE", Suit: "SPADES", Code: "AS"}, "2S": {Value: "2", Suit: "SPADES", Code: "2S"}, "3S": {Value: "3", Suit: "SPADES", Code: "3S"},
	"4S": {Value: "4", Suit: "SPADES", Code: "4S"}, "5S": {Value: "5", Suit: "SPADES", Code: "5S"}, "6S": {Value: "6", Suit: "SPADES", Code: "6S"},
	"7S": {Value: "7", Suit: "SPADES", Code: "7S"}, "8S": {Value: "8", Suit: "SPADES", Code: "8S"}, "9S": {Value: "9", Suit: "SPADES", Code: "9S"},
	"10S": {Value: "10", Suit: "SPADES", Code: "10S"}, "JS": {Value: "JACK", Suit: "SPADES", Code: "JS"}, "QS": {Value: "QUEEN", Suit: "SPADES", Code: "QS"},
	"KS": {Value: "KING", Suit: "SPADES", Code: "KS"},

	"AD": {Value: "ACE", Suit: "DIAMONDS", Code: "AD"}, "2D": {Value: "2", Suit: "DIAMONDS", Code: "2D"}, "3D": {Value: "3", Suit: "DIAMONDS", Code: "3D"},
	"4D": {Value: "4", Suit: "DIAMONDS", Code: "4D"}, "5D": {Value: "5", Suit: "DIAMONDS", Code: "5D"}, "6D": {Value: "6", Suit: "DIAMONDS", Code: "6D"},
	"7D": {Value: "7", Suit: "DIAMONDS", Code: "7D"}, "8D": {Value: "8", Suit: "DIAMONDS", Code: "8D"}, "9D": {Value: "9", Suit: "DIAMONDS", Code: "9D"},
	"10D": {Value: "10", Suit: "DIAMONDS", Code: "10D"}, "JD": {Value: "JACK", Suit: "DIAMONDS", Code: "JD"}, "QD": {Value: "QUEEN", Suit: "DIAMONDS", Code: "QD"},
	"KD": {Value: "KING", Suit: "DIAMONDS", Code: "KD"},

	"AC": {Value: "ACE", Suit: "CLUBS", Code: "AC"}, "2C": {Value: "2", Suit: "CLUBS", Code: "2C"}, "3C": {Value: "3", Suit: "CLUBS", Code: "3C"},
	"4C": {Value: "4", Suit: "CLUBS", Code: "4C"}, "5C": {Value: "5", Suit: "CLUBS", Code: "5C"}, "6C": {Value: "6", Suit: "CLUBS", Code: "6C"},
	"7C": {Value: "7", Suit: "CLUBS", Code: "7C"}, "8C": {Value: "8", Suit: "CLUBS", Code: "8C"}, "9C": {Value: "9", Suit: "CLUBS", Code: "9C"},
	"10C": {Value: "10", Suit: "CLUBS", Code: "10C"}, "JC": {Value: "JACK", Suit: "CLUBS", Code: "JC"}, "QC": {Value: "QUEEN", Suit: "CLUBS", Code: "QC"},
	"KC": {Value: "KING", Suit: "CLUBS", Code: "KC"},

	"AH": {Value: "ACE", Suit: "HEARTS", Code: "AH"}, "2H": {Value: "2", Suit: "HEARTS", Code: "2H"}, "3H": {Value: "3", Suit: "HEARTS", Code: "3H"},
	"4H": {Value: "4", Suit: "HEARTS", Code: "4H"}, "5H": {Value: "5", Suit: "HEARTS", Code: "5H"}, "6H": {Value: "6", Suit: "HEARTS", Code: "6H"},
	"7H": {Value: "7", Suit: "HEARTS", Code: "7H"}, "8H": {Value: "8", Suit: "HEARTS", Code: "8H"}, "9H": {Value: "9", Suit: "HEARTS", Code: "9H"},
	"10H": {Value: "10", Suit: "HEARTS", Code: "10H"}, "JH": {Value: "JACK", Suit: "HEARTS", Code: "JH"}, "QH": {Value: "QUEEN", Suit: "HEARTS", Code: "QH"},
	"KH": {Value: "KING", Suit: "HEARTS", Code: "KH"},
}

var SuitsOrder = map[string]int{"SPADES": 0, "DIAMONDS": 1, "CLUBS": 2, "HEARTS": 3}
var ValuesOrder = map[string]int{"ACE": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "JACK": 11, "QUEEN": 12, "KING": 13}
