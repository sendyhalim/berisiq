package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	_ "github.com/sendyhalim/berisiq/setup"

	"github.com/sendyhalim/berisiq/http/handlers/v1/notifications"
)

func main() {
	router := chi.NewRouter()

	router.Post("/notifications", notifications.SendNotification)

	port := os.Getenv("HTTP_PORT")

	fmt.Printf("Listening on port %s", port)

	http.ListenAndServe(":"+port, router)
}
