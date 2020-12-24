package backend

import (
	database "fpl/database"
	"log"
	"strconv"
)

// getLatestGameweekID will check the DB to get integer value of the latest
// gameweek
func getLatestGameweekID() int {
	log.SetPrefix("GET_LATEST_GAMEWEEK_ID - ")
	db := database.GetDB()
	var currentGameweekStats Event
	if err := db.Where("is_current = ?", true).Find(&currentGameweekStats).Error; err != nil {
		log.Fatal(err)
	}
	log.Printf("Latest gameweek fetch; it is : %d", currentGameweekStats.ID)
	return int(currentGameweekStats.ID)
}

// isGameweekPresentInGameweekStats function checks the DB to find if latest
// game week already has data and this will subsequently lead to data to be
// updated
func isGameweekPresentInGameweekStats(gameweekNumber int) bool {
	log.SetPrefix("IS_GAMEWEEK_PRESENT - ")
	db := database.GetDB()
	type gameweekNumbersStruct struct {
		GameweekNumber int
	}
	var gameweekNumberStruct []gameweekNumbersStruct
	db.Table("gameweek_stats").
		Select("gameweek_number").
		Group("gameweek_number").
		Find(&gameweekNumberStruct)
	log.Printf("Gameweek numbers found: %v\n", gameweekNumberStruct)
	for i := range gameweekNumberStruct {
		if gameweekNumberStruct[i].GameweekNumber == gameweekNumber {
			return true
		}
	}
	return false
}

// UpdatePlayerLatestGameweekData function updates data for each player for
// the latest week. This function will run in the cron job to update the
// database to ensure data is not stale.
func UpdatePlayerLatestGameweekData() {
	log.SetPrefix("UPDATE_PLAYER_LATEST_STATS - ")
	// TODO Cache the latestGameweekID value and if the gameweek is present
	latestGameweekID := getLatestGameweekID()
	isGameweekPresent := isGameweekPresentInGameweekStats(latestGameweekID)
	db := database.GetDB()
	currentWeekStats := FetchGameweekData(latestGameweekID)
	for i := range currentWeekStats.GameweekArray {
		gameweekStat := currentWeekStats.GameweekArray[i].Stats
		gameweekStat.ElementID = currentWeekStats.GameweekArray[i].ID
		gameweekStat.GameweekNumber = latestGameweekID
		primaryKey := strconv.Itoa(gameweekStat.ElementID) +
			"_" + strconv.Itoa(gameweekStat.GameweekNumber)
		gameweekStat.ID = primaryKey
		if isGameweekPresent {
			db.Table("gameweek_stats").Where("ID = ?", primaryKey).
				Delete(GameweekStats{})
		}
		db.Create(&gameweekStat)
		log.Printf("Added data for player %d for game week %d",
			gameweekStat.ElementID, gameweekStat.GameweekNumber)
	}
}

func UpdateGameweekAggregatedData() {

}
