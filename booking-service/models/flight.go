package models

import "time"

type Flight struct {
	FlightID      string    `gorm:"primaryKey" json:"flightId"`
	Name          string    `json:"name"`
	FlightNumber  string    `json:"flightNumber"`
	Price         int       `json:"price"`
	DepartureTime time.Time `json:"departureTime"`
	ArrivalTime   time.Time `json:"arrivalTime"`
	Availability  int       `json:"availability"`
	LocationFrom  string    `json:"locationFrom"`
	LocationTo    string    `json:"locationTo"`
}
