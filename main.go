package main

import (
	"log"
	"net/http"

	_ "EA4WebHook/docs" // Import Swagger docs

	httpSwagger "github.com/swaggo/http-swagger" // Swagger UI package
)

// @Summary Webhook Receiver
// @Description This endpoint receives a webhook and performs an action.
// @Accept json
// @Produce json
// @Param webhook body string true "Webhook Payload"
// @Success 200 {string} string "Webhook received"
// @Failure 400 {string} string "Webhook error"
// @Router /webhook [post]
func handleWebhook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Webhook received"))
}

// @Summary Server Health Check
// @Description This endpoint returns the health status of the server.
// @Success 200 {string} string "Server is running"
// @Router /health [get]
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}

func main() {
	// Swagger UI en la ruta /swagger
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Otros endpoints
	http.HandleFunc("/webhook", handleWebhook)
	http.HandleFunc("/health", healthCheck)

	log.Println("Starting server on http://localhost:8080/webhook")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
