// Package backendfpl is the backend
package backend

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// getHTTPBody creates a HTTP request to the url and returns the body of the
// HTTP page
func getHTTPBody(url string) *[]byte {
	fplClient := http.Client{
		Timeout: time.Second * 2,
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:83.0) Gecko/20100101 Firefox/83.0")
	response, err := fplClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return &body
}

// FetchBootstrapData fetches the players data from the FPL website and element
// stats are overall stats of the player
func FetchBootstrapData() *Bootstrap {
	url := "https://fantasy.premierleague.com/api/bootstrap-static/"
	body := getHTTPBody(url)
	bootstrap := Bootstrap{}
	jsonErr := json.Unmarshal(*body, &bootstrap)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	log.Printf("Bootstrap data fetched successfully\n")
	return &bootstrap
}

// FetchGameweekData fetches data for each player in that specific gameweek
func FetchGameweekData(gameweekNumber int) *Gameweeks {
	apiURL := "https://fantasy.premierleague.com/api/event/{gameweek-number}/live/"
	apiURL = strings.Replace(apiURL, "{gameweek-number}", strconv.Itoa(gameweekNumber), 1)
	body := getHTTPBody(apiURL)
	gameweekStats := Gameweeks{}
	jsonErr := json.Unmarshal(*body, &gameweekStats)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return &gameweekStats
}

// FetchFixturesData fetches fixture data of each player id. The data is only
// forward looking. Player performance history can be downloaded from this API too.
func FetchFixturesData(playerID int) {
	apiURL := "https://fantasy.premierleague.com/api/element-summary/{playerid-number}/"
	apiURL = strings.Replace(apiURL, "{playerid-number}", strconv.Itoa(playerID), 1)
	body := getHTTPBody(apiURL)
	fixture := Fixtures{}
	jsonErr := json.Unmarshal(*body, &fixture)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}
