package item

type Item struct {
	ID       string   `json:"id" bson:"_id,omitempty"`
	Name     string   `json:"name" bson:"name,omitempty"`
	Describe string   `json:"describe" bson:"describe"`
	Category []string `json:"category" bson:"category,omitempty"`
	Price    int      `json:"price" bson:"price,omitempty"`
	Photo    string   `json:"photo" bson:"photo"`
	// laceId  Place
}

type ItemDTO struct {
	Name     string   `json:"name"`
	Describe string   `json:"describe"`
	Category []string `json:"category"`
	Price    int      `json:"price"`
	Photo    string   `json:"photo"`
}
