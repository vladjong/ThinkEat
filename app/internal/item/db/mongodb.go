package db

import (
	"context"
	"fmt"

	"github.com/vladjong/ThinkEat/internal/item"
	"github.com/vladjong/ThinkEat/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, item item.Item) (string, error) {
	d.logger.Debug("Create item")
	result, err := d.collection.InsertOne(ctx, item)
	if err != nil {
		return "", fmt.Errorf("Failed to create item due to errro: %v", err)
	}

	d.logger.Debug("Convert InsertedId to ObjectId")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(item)
	return "", fmt.Errorf("Failde to convert objectId to hex: %s", oid)
}

func (d *db) FindOnly(ctx context.Context, id string) (item.Item, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return item.Item{}, fmt.Errorf("Failed to convert hex to objectId: %s", oid)
	}
	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return item.Item{}, fmt.Errorf("Failed to find one item by id: %s due to error %v", id, err)
	}
	if err = result.Decode(&item.Item{}); err != nil {
		return item.Item{}, fmt.Errorf("Failed to decode user(id:%s) from DB due to error: %v", id, err)
	}
	return item.Item{}, nil
}

func (d *db) Update(ctx context.Context, item item.Item) error {}

func (d *db) Delete(ctx context.Context, id string) error {}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) item.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
