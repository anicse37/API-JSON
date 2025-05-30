package main

import (
	"encoding/json"
	"io"
	"log"
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

func satelliteHandler(w http.ResponseWriter, r *http.Request) {
	var result APIOutput
	url := "https://isro.vercel.app/api/customer_satellites"
	rawdata, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer rawdata.Body.Close()

	body, err := io.ReadAll(rawdata.Body)
	if err != nil {
		http.Error(w, "Failed to Read File", http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &result); err != nil {
		http.Error(w, "failed to decode json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.CustomerSatellite)

}

func main() {
	http.HandleFunc("/data", satelliteHandler)

	fs := http.FileServer(http.Dir("Statics"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":8082", nil))
}
