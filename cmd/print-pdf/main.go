package main

import (
	"log"

	config "github.com/Nyralei/print-pdf/internal/config"
	internalui "github.com/Nyralei/print-pdf/internal/ui"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	loc, err := internalui.LoadLocalization(cfg.Language)
	if err != nil {
		log.Fatalf("Error loading localization: %v", err)
	}

	err = ui.Main(func() {
		internalui.SetupUI(cfg, loc)
	})
	if err != nil {
		log.Fatalf("Error starting UI: %v", err)
	}
}
