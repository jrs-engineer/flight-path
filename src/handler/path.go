package handler

import (
	"encoding/json"
	"net/http"

	"github.com/C001-developer/flight-path/src/logic"
)

const (
	// ErrInvalidJSONData is a message that is returned when the request body contains invalid JSON data.
	ErrInvalidJSONData = "invalid JSON data"
)

// FlightPathHandler is a handler function for POST requests to "/calculate" endpoint.
// It expects a JSON array of flight paths in the request body.
// It returns an exact path [source, target].
func FlightPathHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON data from the request body into a slice of Flight structs
	var flights [][2]string
	err := json.NewDecoder(r.Body).Decode(&flights)
	if err != nil {
		http.Error(w, ErrInvalidJSONData, http.StatusBadRequest)
		return
	}

	res, err := logic.GetSinglePath(flights)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send a response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}
