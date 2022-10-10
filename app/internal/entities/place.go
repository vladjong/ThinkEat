package entities

import "time"

type Place struct {
	ID       string   `json:"id" bson:"_id,omitempty"`
	Name     string   `json:"name" bson:"name,omitempty"`
	Describe string   `json:"describe" bson:"describe,omitempty"`
	URL      string   `json:"url" bson:"url"`
	PhotoId  string   `json:"photoId" bson:"photoId"`
	Category Category `json:"category" bson:"category"`
	Contacts Contacts `json:"contacts" bson:"contacts"`
}

type Category struct {
	Class   string `json:"class" bson:"class,omitempty"`
	Kitchen string `json:"kitchen" bson:"kitchen,omitempty"`
}

type Contacts struct {
	Address  Address  `json:"adress" bson:"adress,omitempty"`
	TimeWork TimeWork `json:"time_work" bson:"time_work"`
}

type Address struct {
	City   string `json:"city" bson:"city,omitempty"`
	Street string `json:"street" bson:"street,omitempty"`
}

type TimeWork struct {
	Opening time.Time `json:"opening" bson:"opening,omitempty"`
	Closure time.Time `json:"closure" bson:"closure,omitempty"`
	Weekend string    `json:"weekend" bson:"weekend,omitempty"`
}

func NewCategory(class, kitchen string) *Category {
	return &Category{
		Class:   class,
		Kitchen: kitchen,
	}
}

func NewContacts(city, street, weekend string, opening, closure time.Time) *Contacts {
	return &Contacts{
		Address: Address{
			City:   city,
			Street: street,
		},
		TimeWork: TimeWork{
			Opening: opening,
			Closure: closure,
			Weekend: weekend,
		},
	}
}

func NewPlace(id, name, describe, url, photoId string, category Category, contacts Contacts) *Place {
	return &Place{
		ID:       id,
		Name:     name,
		Describe: describe,
		URL:      url,
		PhotoId:  photoId,
		Category: category,
		Contacts: contacts,
	}
}
