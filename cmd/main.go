package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rvmzx/service-booking/api"
	"github.com/rvmzx/service-booking/internal/handler"
	"github.com/rvmzx/service-booking/internal/storage"
)

func main() {
	ctx := context.Background()

	db, err := storage.SetupDatabase(ctx)
	if err != nil {
		panic(err)
	}

	bookingManager := api.NewBookingManager(db)
	router := handler.NewRouter(bookingManager)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
