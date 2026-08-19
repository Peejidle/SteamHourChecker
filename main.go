package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"bufio"
	"os"
)

type SteamJson struct {
	Response					SteamData 	`json:"response"`
}

type SteamData struct {
	GameCount					int					`json:"game_count"`
	Games							[]GameData	`json:"games"`
}

type GameData struct {
	Name								string 		`json:"name"`
	AppID								int				`json:"appid"`
	PlaytimeForever			int				`json:"playtime_forever"`
	Playtime2Weeks			int				`json:"playtime_2weeks"`
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	rawData, err := getData()
	if err != nil {
		fmt.Println("Error: Failed to Fetch Steam Data")
		return
	}

	parsedData, err := unmarshal(rawData)
	if err != nil {
		fmt.Println("Error: Failed to Parse Steam Data")
		return
	}

	fmt.Println("Welcome to Steam Hour Checker.","\nFirst enter your personal steam API key.(NEVER SHARE THIS)")
	scanner.Scan()
	usrAPIKey := scanner.Text()

	fmt.Prinln("Second, enter your steamID(The numbers in your profile link without the rest of the link. just the numbers)\n")
	scanner.Scan()
	usrSteamID := scanner.Text()

	fmt.Println("Now what game do you want to check?\n")
	scanner.Scan()
	userGame := scanner.Text()



	for _, Games := range parsedData.Response.Games {
		if Games.Name == userGame {

			foreverHours := Games.PlaytimeForever / 60
			weeksHours := Games.Playtime2Weeks / 60

			fmt.Printf("Name: %v\n", Games.Name)
			fmt.Printf("Total Playtime: %v\n", foreverHours)
			fmt.Printf("Last 2 Weeks: %v\n", weeksHours)
			return 
		}
	}
	fmt.Println("No game found, Check for typos or try again")
}

func getData() ([]byte, error){
	resp, err := http.Get("PLACEHOLDER") // used to be hard coded in. Working on changing this.
	if err != nil {
		fmt.Println("Error fetching URL", err)
		return nil, err
	}
	defer resp.Body.Close()

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading data", err)
		return nil, err
	}
	return rawData, nil
}

func unmarshal(rawData []byte) (SteamJson, error) {
	var parsedData SteamJson
	err := json.Unmarshal(rawData, &parsedData)
	if err != nil {
		fmt.Println("Error Unmarsaling data", err)
		return parsedData, err
	}
	return parsedData, nil
}
