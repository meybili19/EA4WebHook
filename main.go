package main

import (
	"log"
	"net/http"

	_ "EA4WebHook/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @Summary Estado del servidor
// @Description Este endpoint devuelve el estado del servidor.
// @Success 200 {string} string "Servidor en funcionamiento"
// @Router /health [get]
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook server is running"))
}

// @Summary Recibe el webhook
// @Description Este endpoint recibe un webhook y ejecuta una acción.
// @Accept json
// @Produce json
// @Param webhook body string true "Webhook Payload"
// @Success 200 {string} string "Webhook recibido"
// @Failure 400 {string} string "Error en el webhook"
// @Router /webhook [post]
func handleWebhook(w http.ResponseWriter, r *http.Request) {
	// Lógica del webhook
	w.Write([]byte("Webhook recibido"))
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
