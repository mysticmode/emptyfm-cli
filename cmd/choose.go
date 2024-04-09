package cmd

import (
	"log"

	"github.com/pterm/pterm"
)

func Chooser(options ...string) string {
	selectedOption, err := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	if err != nil {
		log.Fatal(err)
	}
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))

	return selectedOption
}
