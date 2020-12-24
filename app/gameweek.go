package app

import (
	"encoding/json"
	backend "fpl/backend"
	database "fpl/database"
	"net/http"
)

func GetGameweekAggregated(w http.ResponseWriter, r *http.Request) {
	var events []backend.Event
	db := database.GetDB()
	w.Header().Set("Content-Type", "application/json")
	// TODO: Add caching here :Caching:
	db.Table("events").Find(&events)
	json.NewEncoder(w).Encode(events)
}
