package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
*
* This program uses an API to get some JSON data and convert it
* and assign it to structure to use efficiently.
*
* Toconvert the file it uses Unmarshal function in "encoding/json" package
*
* Also, Data canbe accessed individually using structure
*
 */

type APIOutput struct {
	CustomerSatellite []customer_satellites `json:"customer_satellites"`
}
type customer_satellites struct {
	Id          string `json:"id"`
	Country     string `json:"country"`
	Launch_date string `json:"launch_date"`
	Launcher    string `json:"launcher"`
}

func main() {
	var result APIOutput
	url := "https://isro.vercel.app/api/customer_satellites"
	rawData, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occured while loading api")
		return
	}
	defer rawData.Body.Close()
	body, err := io.ReadAll(rawData.Body)
	if err != nil {
		fmt.Println("Error occured while loading data")
		return
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error occured while decording json")
		return
	}
	// fmt.Println(result.CustomerSatellite)
	for _, individualData := range result.CustomerSatellite {
		fmt.Println("----------------------")
		fmt.Println("ID:", individualData.Id)
		fmt.Println("Country:", individualData.Country)
		fmt.Println("Launch Date:", individualData.Launch_date)
		fmt.Println("Mission Type:", individualData.Launcher)
	}

}
