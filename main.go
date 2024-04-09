package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mysticmode/emptyfm-cli/lib"
	"github.com/pterm/pterm"
)

type FMStation struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func main() {
	var langStore lib.Store
	err := langStore.Parser("./iso-639-lang.json")
	if err != nil {
		log.Fatalf("error occured when parsing language JSON file: %v", err)
	}

	var countryStore lib.Store
	err = countryStore.Parser("./iso-3166-1-alpha2-countrycode.json")
	if err != nil {
		log.Fatalf("Error occured when parsing language JSON file: %v", err)
	}

	// selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(lngs).Show()
	// pterm.Info.Printfln("Selected option: %s:%s", pterm.Green(selectedOption), pterm.Yellow(langMap[selectedOption]))

	fmt.Println(langStore.Mapper["Tamil"])
	fmt.Println(countryStore.Mapper["India"])

	configPath := lib.ConfigPath("emptyfm-cli", "config.yml")

	err = os.MkdirAll(filepath.Dir(configPath), 0700)
	if err == nil {
		err = os.WriteFile(configPath, nil, 0600)
	}
	if err != nil {
		log.Fatalf("error saving config changes: %v", err)
	}

	radioBrowser := "https://de1.api.radio-browser.info/json/stations/search?" + "countrycode=" + countryStore.Mapper["India"] + "&language=" + langStore.Mapper["Tamil"] + "&hidebroken=true&order=clickcount&reverse=true"
	response, err := http.Get(radioBrowser)
	if err != nil {
		log.Fatalf("Error on GET request to radio-browser: %v", err)
	}

	var fmStations []FMStation
	if err := json.NewDecoder(response.Body).Decode(&fmStations); err != nil {
		log.Fatalf("Error on parsing the response: %v", err)
	}

	pterm.Info.Printfln("Playing Kodai FM..")

	cmd := exec.Command("ffplay", "-nodisp", "-loglevel", "error", "https://air.pc.cdn.bitgravity.com/air/live/pbaudio051/playlist.m3u8")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
