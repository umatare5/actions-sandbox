package testdata

import (
	"log"
	"net/http"
	"os"
)

//lint:ignore U1000 main is intentionally unused in a non-main package (example server)
//nolint:staticcheck // U1000: main is intentionally unused in a non-main package
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/health", HealthHandler)

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
