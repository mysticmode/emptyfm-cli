package main

import "github.com/mysticmode/emptyfm-cli/cmd"

type FMStation struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func main() {

	cmd.Execute()

	// fmt.Println(langStore.Mapper["Tamil"])
	// fmt.Println(countryStore.Mapper["India"])

	// radioBrowser := "https://de1.api.radio-browser.info/json/stations/search?" + "countrycode=" + countryStore.Mapper["India"] + "&language=" + langStore.Mapper["Tamil"] + "&hidebroken=true&order=clickcount&reverse=true"
	// response, err := http.Get(radioBrowser)
	// if err != nil {
	// 	log.Fatalf("Error on GET request to radio-browser: %v", err)
	// }

	// var fmStations []FMStation
	// if err := json.NewDecoder(response.Body).Decode(&fmStations); err != nil {
	// 	log.Fatalf("Error on parsing the response: %v", err)
	// }

	// pterm.Info.Printfln("Playing Kodai FM..")

	// cmd := exec.Command("ffplay", "-nodisp", "-loglevel", "error", "https://air.pc.cdn.bitgravity.com/air/live/pbaudio051/playlist.m3u8")
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
