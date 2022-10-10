package entities

type Item struct {
	ID       string   `json:"id" bson:"_id,omitempty"`
	Name     string   `json:"name" bson:"name,omitempty"`
	Describe string   `json:"describe" bson:"describe"`
	Category []string `json:"category" bson:"category,omitempty"`
	Price    float64  `json:"price" bson:"price,omitempty"`
	Weight   float64  `json:"weight" bson:"weight"`
	Photo    string   `json:"photo" bson:"photo"`
	PlaceId  string   `json:"place_id" bson:"place_id"`
}

func NewItem(id, name, describe, photo, placeId string, category []string, price, weight float64) *Item {
	return &Item{
		ID:       id,
		Name:     name,
		Describe: describe,
		Category: category,
		Price:    price,
		Weight:   weight,
		Photo:    photo,
		PlaceId:  placeId,
	}
}
