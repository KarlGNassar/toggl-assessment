package api

import (
	"fmt"
)

const (
	ParsingFailed     = "PARSING_FAILED"
	InsufficientCards = "INSUFFICIENT_CARDS"
	InvalidParam      = "INVALID_PARAM"
)

type ErrDeckNotFound struct {
	Id string
}

func (e *ErrDeckNotFound) Error() string {
	return fmt.Sprintf("deck with id %s not found", e.Id)
}

type ErrInsufficientCards struct {
	Requested, Available int
}

func (e *ErrInsufficientCards) Error() string {
	return fmt.Sprintf("not enough cards in the deck: requested %d, available %d", e.Requested, e.Available)
}
