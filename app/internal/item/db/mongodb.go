package db

import (
	"context"
	"errors"
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
		return "", fmt.Errorf("failed to create item due to errro: %v", err)
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(item)
	return "", fmt.Errorf("failde to convert objectId to hex: %s", oid)
}

func (d *db) FindID(ctx context.Context, id string) (item item.Item, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return item, fmt.Errorf("failed to convert hex to objectId: %s", oid)
	}
	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			//TODO
			return item, fmt.Errorf("error not found documents")
		}
		return item, fmt.Errorf("failed to find one item by id: %s due to error %v", id, err)
	}
	if err = result.Decode(&item); err != nil {
		return item, fmt.Errorf("failed to decode item(id:%s) from DB due to error: %v", id, err)
	}
	return item, nil
}

func (d *db) Update(ctx context.Context, item item.Item) error {
	objectID, err := primitive.ObjectIDFromHex(item.ID)
	if err != nil {
		return fmt.Errorf("failed to convert item ID to ObjectID ID=%s", item.ID)
	}
	filter := bson.M{"_id": objectID}
	itemBytes, err := bson.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal item errror: %v", err)
	}
	var updateItemObj bson.M
	err = bson.Unmarshal(itemBytes, &updateItemObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal item bytes error: %v", err)
	}
	delete(updateItemObj, "_id")
	update := bson.M{
		"$set": updateItemObj,
	}
	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update item filter error: %v", err)
	}
	if result.MatchedCount == 0 {
		//TODO
		return fmt.Errorf("failed not found item")
	}
	d.logger.Tracef("Matched %d documents and Modified %d documents", result.MatchedCount, result.ModifiedCount)
	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert item ID to ObjectID ID=%s", id)
	}
	filter := bson.M{"_id": objectID}
	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute filter error: %v", err)
	}
	d.logger.Tracef("Deleted %d documents", result.DeletedCount)
	return nil
}

func (d *db) FindAll(ctx context.Context) (items []item.Item, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return items, fmt.Errorf("failed to find all itens due to error: %v", err)
	}
	if err = cursor.All(ctx, &items); err != nil {
		return items, fmt.Errorf("failed to read all documents error: %v", err)
	}
	return items, nil
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) item.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
