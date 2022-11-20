package entities

import "time"

type Place struct {
	ID       int      `json:"id" db:"id"`
	Name     string   `json:"name" db:"name,omitempty"`
	Describe string   `json:"describe" db:"describe"`
	URL      string   `json:"url" db:"url"`
	Photo    string   `json:"photo" db:"photo"`
	Class    string   `json:"class" db:"class"`
	Contacts Contacts `json:"contacts" db:"contacts"`
}

type Contacts struct {
	City         string `json:"city" db:"city"`
	Street       string `json:"street" db:"street"`
	WorkSchedule WorkSchedule
}

type WorkSchedule struct {
	BegginTime time.Time `json:"beggin_time" db:"beggin_time"`
	EndTime    time.Time `json:"end_time" db:"end_time"`
	Weekend    int       `json:"weekend" db:"weekend"`
	IsDayOff   bool      `json:"is_day_off" db:"is_day_off"`
}
