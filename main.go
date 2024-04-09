package main

import (
	"fmt"
	"log"

	"github.com/mysticmode/emptyfm-cli/lib"
)

func main() {
	var langStore lib.Store
	err := langStore.Parser("./iso-639-lang.json")
	if err != nil {
		log.Fatalf("Error occured when parsing language JSON file: %v", err)
	}

	var countryStore lib.Store
	err = countryStore.Parser("./iso-3166-1-alpha2-countrycode.json")
	if err != nil {
		log.Fatalf("Error occured when parsing language JSON file: %v", err)
	}

	// selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(lngs).Show()
	// pterm.Info.Printfln("Selected option: %s:%s", pterm.Green(selectedOption), pterm.Yellow(langMap[selectedOption]))

	fmt.Println(langStore.Text)
	fmt.Println()
	fmt.Println(countryStore.Text)
	fmt.Println(lib.ConfigPath("emptyfm-cli"))
}
