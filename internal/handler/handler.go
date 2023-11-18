package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rvmzx/service-booking/api"
	"github.com/rvmzx/service-booking/internal/storage"
)

func NewRouter(bm *api.BookingManager) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/bookings", handleBookings(bm)).Methods("GET")
	router.HandleFunc("/book", handleBook(bm)).Methods("POST")

	router.HandleFunc("/services", handleServices(bm)).Methods("GET")
	router.HandleFunc("/service", handleService(bm)).Methods("POST")

	return router

}

func handleBookings(bm *api.BookingManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookings, err := bm.GetAllBookings(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseJSON, err := json.Marshal(bookings)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
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

		var booking storage.Booking
		err = json.Unmarshal(body, &booking)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		err = bm.NewBooking(r.Context(), &booking)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("booking added"))
	}
}

func handleServices(bm *api.BookingManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookings, err := bm.GetAllServices(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseJSON, err := json.Marshal(bookings)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
	}
}

func handleService(bm *api.BookingManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		fmt.Println("Received POST request with body:", string(body))

		var service storage.Service
		err = json.Unmarshal(body, &service)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		err = bm.NewService(r.Context(), &service)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("service added"))
	}
}
