package testdata

import (
	"fmt"
	"net/http"
)

// HelloHandler handles HTTP requests to the hello endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello, %s!\n", name)
}

// HealthHandler handles health check requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
