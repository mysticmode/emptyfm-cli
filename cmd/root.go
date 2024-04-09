package cmd

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pterm/pterm"
)

type FMStation struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

var StationNames []string

func Execute() {
	selectedCountry := Chooser(CountryStore.Text...)
	selectedLanguage := Chooser(LangStore.Text...)

	radioBrowser := "https://de1.api.radio-browser.info/json/stations/search?" + "countrycode=" + CountryStore.Mapper[selectedCountry] + "&language=" + LangStore.Mapper[selectedLanguage] + "&hidebroken=true&order=clickcount&reverse=true"

	response, err := http.Get(radioBrowser)
	if err != nil {
		log.Fatalf("Error on GET request to radio-browser: %v", err)
	}

	var fmStations []FMStation
	if err := json.NewDecoder(response.Body).Decode(&fmStations); err != nil {
		log.Fatalf("Error on parsing the response: %v", err)
	}

	stations := make(map[string]string)

	for i := 0; i < len(fmStations); i++ {
		stations[fmStations[i].Name] = fmStations[i].URL
		StationNames = append(StationNames, fmStations[i].Name)
	}

	selectedFMStation := Chooser(StationNames...)

	url := stations[selectedFMStation]

	pterm.Info.Printfln("Playing " + selectedFMStation)

	Play(url)
}
