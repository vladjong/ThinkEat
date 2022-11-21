package entities

type PlacePost struct {
	ID       int         `json:"id" db:"id"`
	Name     string      `json:"name" db:"name,omitempty"`
	Describe string      `json:"describe" db:"describe"`
	URL      string      `json:"url" db:"url"`
	Photo    string      `json:"photo" db:"photo"`
	Class    string      `json:"class" db:"class"`
	Contact  ContactPost `json:"contact"`
}

type ContactPost struct {
	City         string         `json:"city" db:"city"`
	Street       string         `json:"street" db:"street"`
	WorkSchedule []WorkSchedule `json:"work_schedule"`
}

type WorkSchedule struct {
	BegginTime string `json:"beggin_time" db:"beggin_time"`
	EndTime    string `json:"end_time" db:"end_time"`
	Weekday    int    `json:"weekday" db:"weekday"`
	IsDayOff   bool   `json:"is_day_off" db:"is_day_off"`
}

type Place struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name,omitempty"`
	Describe string `json:"describe" db:"describe"`
	URL      string `json:"url" db:"url"`
	Photo    string `json:"photo" db:"photo"`
	Class    string `json:"class" db:"class"`
	Contact
}

type Contact struct {
	City   string `json:"city" db:"city"`
	Street string `json:"street" db:"street"`
}
