package models

type Booking struct {
	ID       int    `json:"id"`
	CourtID  int    `json:"court_id"`
	Date     string `json:"date"`
	TimeSlot string `json:"time_slot"`
	Status   string `json:"status"`
}
