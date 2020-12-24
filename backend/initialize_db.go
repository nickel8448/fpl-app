package backend

import (
	database "fpl/database"
	"log"
	"strconv"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func initializeElementsDB(db *gorm.DB) {
	log.SetPrefix("ELEMENTS_DB_INITIALIZE - ")
	db.AutoMigrate(&Elements{})
	bootstrap := FetchBootstrapData()
	validate := validator.New()
	var allPlayers []Elements
	for i := range bootstrap.Players {
		err := validate.Struct(bootstrap.Players[i])
		allPlayers = append(allPlayers, bootstrap.Players[i])
		if err != nil {
			log.Fatalf(err.Error())
			return
		}
	}
	log.Printf("Data fetched from FPL API")
	db.Create(&allPlayers)
	log.Printf("Initialized Elements DB")
}

func initializeGameweekEvents(db *gorm.DB) {
	log.SetPrefix("GAMEWEEK_EVENTS_INITIALIZE - ")
	db.AutoMigrate(&Event{})
	bootstrap := FetchBootstrapData()
	for i := range bootstrap.Events {
		db.Create(&bootstrap.Events[i])
	}
	log.Printf("Initialized Gameweeks DB")
}

// initalizeGameweekForPlayersDB
func initializeGameweekForPlayersDB(db *gorm.DB) {
	log.SetPrefix("GAMEWEEK_FOR_PLAYERS_DB_INITIALIZE - ")
	var eventStats []Event
	db.Where("finished = ? AND is_current = ?", true, false).Find(&eventStats)
	// TODO Add error handling for record
	numberOfGameweeksPassed := len(eventStats)
	log.Printf("Number of gameweeks passed: %d\n", numberOfGameweeksPassed)
	db.AutoMigrate(&GameweekStats{})
	for i := 1; i <= numberOfGameweeksPassed; i++ {
		gameweekStats := FetchGameweekData(i)
		for j := range gameweekStats.GameweekArray {
			gameweekStat := gameweekStats.GameweekArray[j].Stats
			gameweekStat.ElementID = gameweekStats.GameweekArray[j].ID
			gameweekStat.GameweekNumber = i
			primaryKey := strconv.Itoa(gameweekStat.ElementID) +
				"_" + strconv.Itoa(gameweekStat.GameweekNumber)
			gameweekStat.ID = primaryKey
			db.Create(&gameweekStat)
			log.Printf("Added data for player %d for game week %d",
				gameweekStat.ElementID, gameweekStat.GameweekNumber)
		}
	}
}

// InitalizeDB function initializes the DB for Elements (Players),
// GameweekEvents (Aggregated data of each week) and GameweekForPlayersDB
// (Individual player data for each week)
func InitalizeDB() {
	db := database.GetDB()
	initializeElementsDB(db)
	// initializeGameweekEvents(db)
	// initializeGameweekForPlayersDB(db)
}
