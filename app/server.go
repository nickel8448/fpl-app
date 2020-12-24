package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Server() {
	router := mux.NewRouter()
	// Router handlers / End points
	router.HandleFunc("/api/player/{player_id}", GetPlayer).Methods("GET")
	router.HandleFunc("/api/player_gameweek/{player_id}-{gameweek_id}",
		GetPlayerGameweek).Methods("GET")
	router.HandleFunc("/api/gameweek_aggregated", GetGameweekAggregated).
		Methods("GET")
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}
