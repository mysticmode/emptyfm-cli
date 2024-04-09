package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mysticmode/emptyfm-cli/lib"
)

func init() {
	configPath := lib.ConfigPath("emptyfm-cli", "config.yml")

	err := os.MkdirAll(filepath.Dir(configPath), 0700)
	if err == nil {
		err = os.WriteFile(configPath, nil, 0600)
	}
	if err != nil {
		log.Fatalf("error saving config changes: %v", err)
	}
}

func ParseLibrary(lang, countrycode []byte) (langStore lib.Store, countryStore lib.Store) {
	err := langStore.Parser(lang)
	if err != nil {
		log.Fatalf("error occured when parsing language JSON file: %v", err)
	}

	err = countryStore.Parser(countrycode)
	if err != nil {
		log.Fatalf("Error occured when parsing country JSON file: %v", err)
	}

	return
}
