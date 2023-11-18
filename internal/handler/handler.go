package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rvmzx/service-booking/api"
)

func NewRouter(bm *api.BookingManager) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/bookings", handleBookings(bm)).Methods("GET")
	router.HandleFunc("/book", handleBook(bm)).Methods("POST")
	return router

}

func handleBookings(bm *api.BookingManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GET request received"))
	}
}

func handleBook(bm *api.BookingManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		fmt.Println("Received POST request with body:", string(body))

		bm.NewBooking(r.Context(), body)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("POST request received successfully"))
	}
}
