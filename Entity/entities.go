package entity

import "date"


//Flight represents the information about the flight
type Flight struct {
	ID       int
	Status string 
	Destination_id  int
	Plane_id    int 
	Destination_id date.date
}
