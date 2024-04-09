package main

import (
	"embed"
	"log"

	"github.com/mysticmode/emptyfm-cli/cmd"
)

//go:embed iso-639-lang.json
var isoLang embed.FS

//go:embed iso-3166-1-alpha2-countrycode.json
var isoCountryCode embed.FS

func main() {
	langContent, err := isoLang.ReadFile("iso-639-lang.json")
	if err != nil {
		log.Fatal(err)
	}

	countryCodeContent, err := isoCountryCode.ReadFile("iso-3166-1-alpha2-countrycode.json")
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute(cmd.ParseLibrary(langContent, countryCodeContent))
}
