// endpoints.go
package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheckHandler : Healthcheck
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "hello world kanchan kumar")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)
	log.Println("server start on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
