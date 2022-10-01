package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database string, authDB string) (*mongo.Database, error) {
	isAuth := true
	if authDB != "" {
		isAuth = false
	}
	mongoDBURL := fmt.Sprintf("mongodb://%s:%s", host, port)
	clientOpts := options.Client().ApplyURI(mongoDBURL)
	if isAuth {
		credential := options.Credential{
			Username: username,
			Password: password,
		}
		clientOpts.SetAuth(credential)
	}
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB: %v", err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB: %v", err)
	}
	return client.Database(database), nil
}
