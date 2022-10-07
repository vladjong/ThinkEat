package mongoDB

import (
	"context"
	"errors"
	"fmt"

	"github.com/vladjong/ThinkEat/internal/entities"
	"github.com/vladjong/ThinkEat/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type itemStorage struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) *itemStorage {
	return &itemStorage{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

func (d *itemStorage) Create(ctx context.Context, item *entities.Item) (string, error) {
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

func (d *itemStorage) GetID(ctx context.Context, id string) (item *entities.Item, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return item, fmt.Errorf("failed to convert hex to objectId: %s", oid)
	}
	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return item, fmt.Errorf("error not found documents")
		}
		return item, fmt.Errorf("failed to find one item by id: %s due to error %v", id, err)
	}
	if err = result.Decode(&item); err != nil {
		return item, fmt.Errorf("failed to decode item(id:%s) from DB due to error: %v", id, err)
	}
	return item, nil
}

func (d *itemStorage) GetName(ctx context.Context, name string) (items []*entities.Item, err error) {
	filter := bson.M{"name": bson.M{"$regex": name}}
	cursor, err := d.collection.Find(ctx, filter)
	if err != nil {
		return items, fmt.Errorf("failed to find all itens due to error: %v", err)
	}
	if err = cursor.All(ctx, &items); err != nil {
		return items, fmt.Errorf("failed to read all documents error: %v", err)
	}
	return items, nil
}

func (d *itemStorage) Update(ctx context.Context, item *entities.Item) error {
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
		return fmt.Errorf("failed not found item")
	}
	d.logger.Tracef("Matched %d documents and Modified %d documents", result.MatchedCount, result.ModifiedCount)
	return nil
}

func (d *itemStorage) Delete(ctx context.Context, id string) error {
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

func (d *itemStorage) GetAll(ctx context.Context) (items []*entities.Item, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return items, fmt.Errorf("failed to find all itens due to error: %v", err)
	}
	if err = cursor.All(ctx, &items); err != nil {
		return items, fmt.Errorf("failed to read all documents error: %v", err)
	}
	return items, nil
}
