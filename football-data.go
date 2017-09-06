package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type (
	// Fixtures - Struct to hold a match week of fixtures
	Fixtures struct {
		Count    int       `json:"count"`
		Fixtures []Fixture `json:"fixtures"`
	}

	// Fixture - Struct to hold the details of an individual fixture
	Fixture struct {
		fixtureID     int    `json:"id"`
		CompetitionID int    `json:"competitionId"`
		Date          string `json:"date"`
		Status        string `json:"status"`
		Matchday      int    `json:"matchday"`
		HomeTeamName  string `json:"homeTeamName"`
		HomeTeamID    int    `json:"homeTeamId"`
		AwayTeamName  string `json:"awayTeamName"`
		AwayTeamID    int    `json:"awayTeamId"`
		Result        Result `json:"result"`
		Odds          int    `json:"odds"`
	}

	// Result - Struct to hold the match result for a fixture
	Result struct {
		GoalsHomeTeam int `json:"goalsHomeTeam"`
		GoalsAwayTeam int `json:"goalsAwayTeam"`
	}
)

func main() {

	var req *http.Request
	var err error
	var expectedResponse int

	url := "http://api.football-data.org/v1/competitions/398/fixtures?matchday=8"

	req, err = http.NewRequest("GET", url, nil)
	expectedResponse = http.StatusOK
	req.Header.Set("Content-Type", "application/json")
	apikey := os.Getenv("FOOTBALL_API_KEY")
	if apikey != "" {
		req.Header.Set("X-Access-Token", apikey)
	}
	req.Header.Set("X-Response-Control", "minified")

	// Make the actual request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error is " + err.Error())
	}
	defer resp.Body.Close()

	// Make sure we get the correct response.
	if resp.StatusCode != expectedResponse {
		fmt.Println("Unexpected response from API, expected %d got: %d", expectedResponse, resp.StatusCode)
	}

	// Define a new fixtures object
	f := Fixtures{}

	// Decode body in interface{} here, because the Body.Close() means decode doesn't work after stream is closed.
	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		fmt.Println("Invalid JSON in HTTP Response from API %s", err.Error())
	}

	for _, fixture := range f.Fixtures {
		fmt.Printf("%s %d - %d %s\n", fixture.HomeTeamName, fixture.Result.GoalsHomeTeam, fixture.Result.GoalsAwayTeam, fixture.AwayTeamName)
	}

}
