package cmd

import (
	"log"
	"os/exec"
)

func Play(url string) {
	cmd := exec.Command("ffplay", "-nodisp", "-loglevel", "error", url)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
