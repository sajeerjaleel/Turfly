package handlers

import (
	"net/http"
)

func GetBookingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of bookings"))
}
