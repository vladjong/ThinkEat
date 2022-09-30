package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database string) (*mongo.Database, error) {
	mongoDBURL := fmt.Sprintf("mongo://%s:%s", host, port)
	credential := options.Credential{
		Username: username,
		Password: password,
	}
	clientOpts := options.Client().ApplyURI(mongoDBURL).
		SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to mongoDB: %v", err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Failed to connect to mongoDB: %v", err)
	}
	return client.Database(database), nil
}
