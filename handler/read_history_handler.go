package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/comet-backend/repository"
	"github.com/gorilla/mux"
)

func ReadHistoryConfiguration(w http.ResponseWriter, r *http.Request) {
	historyRepo := repository.NewConfigurationRepository()
	params := mux.Vars(r)
	header := r.Header.Get("Authorization")

	history := historyRepo.ReadHistory(params["appName"], params["namespaceName"], header)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", header)
	json.NewEncoder(w).Encode(history)
}
