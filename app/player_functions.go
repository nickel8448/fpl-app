// Package app provides API functions
package app

import (
	"encoding/json"
	backend "fpl/backend"
	database "fpl/database"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	var player backend.Elements
	db := database.GetDB()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// TODO: Add caching here :Caching:
	playerID := params["player_id"]
	db.Debug().Where("id = ?", playerID).Find(&player)
	json.NewEncoder(w).Encode(player)
}

func GetPlayerGameweek(w http.ResponseWriter, r *http.Request) {
	var playerGameweek backend.GameweekStats
	db := database.GetDB()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	playerIDGameweekID := params["player_id"] + "_" + params["gameweek_id"]
	db.Debug().Where("id = ?", playerIDGameweekID).Find(&playerGameweek)
	json.NewEncoder(w).Encode(playerGameweek)
}
