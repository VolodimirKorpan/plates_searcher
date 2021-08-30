package models

import "time"

type Plate struct {
	ID       int       `json:"id"`
	DateTime time.Time `json:"date_time"`
	Camera   string    `json:"camera"`
	Plate    string    `json:"plate"`
	Image    string    `json:"image"`
}

type SearchPlate struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
	NumPlate  string `json:"num_plate"`
}
