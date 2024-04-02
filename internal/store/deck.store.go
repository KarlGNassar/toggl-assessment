package store

import (
	"context"

	"github.com/KarlGNassar/toggl-assessment/internal/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeckStore interface {
	CreateDeck(ctx context.Context, deck model.Deck) error
	GetDeck(ctx context.Context, id uuid.UUID) (model.Deck, error)
	UpdateDeck(ctx context.Context, id uuid.UUID, deck model.Deck) error
}

type MongoDeckStore struct {
	db *mongo.Collection
}

func NewMongoDeckStore(client *mongo.Client, dbName string, collectionName string) *MongoDeckStore {
	collection := client.Database(dbName).Collection(collectionName)
	return &MongoDeckStore{db: collection}
}

func (s *MongoDeckStore) CreateDeck(ctx context.Context, deck model.Deck) error {
	_, err := s.db.InsertOne(ctx, deck)
	return err
}

func (s *MongoDeckStore) GetDeck(ctx context.Context, id uuid.UUID) (model.Deck, error) {
	var deck model.Deck
	filter := bson.M{"id": id}
	err := s.db.FindOne(ctx, filter).Decode(&deck)
	return deck, err
}

func (s *MongoDeckStore) UpdateDeck(ctx context.Context, id uuid.UUID, deck model.Deck) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"cards": deck.Cards, "remaining": deck.Remaining}}
	_, err := s.db.UpdateOne(ctx, filter, update)
	return err
}
