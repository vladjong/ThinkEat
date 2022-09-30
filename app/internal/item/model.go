package item

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID       primitive.ObjectID `json:"id" bson:"_id, omitempy"`
	Name     string             `json:"name" bson:"name, omitempy"`
	Describe string             `json:"describe" bson:"describe"`
	Category []string           `json:"category" bson:"category, omitempy"`
	Price    int                `json:"price" bson:"price, omitempy"`
	Photo    string             `json:"photo" bson:"photo"`
	// laceId  Place
}

type ItemDTO struct {
	Name     string   `json:"name"`
	Describe string   `json:"describe"`
	Category []string `json:"category"`
	Price    int      `json:"price"`
	Photo    string   `json:"photo"`
}
