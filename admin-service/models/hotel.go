package models

type Hotel struct {
	HotelID      string `gorm:"primaryKey" json:"hotelId"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Availability int    `json:"availability"`
	Location     string `json:"location"`
}
