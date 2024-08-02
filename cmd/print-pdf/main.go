package main

import (
	"log"
	"os"
	"path/filepath"

	config "github.com/Nyralei/print-pdf/internal/config"
	internalui "github.com/Nyralei/print-pdf/internal/ui"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

const configDir = ".config/print-pdf"
const configFileName = "config.json"

func main() {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home directory: %v", err)
	}

	// Create the full path to the config directory
	configDirPath := filepath.Join(homeDir, configDir)

	// Create the config directory if it doesn't exist
	err = os.MkdirAll(configDirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating config directory: %v", err)
	}

	// Full path to the config file
	configFilePath := filepath.Join(configDirPath, configFileName)

	// Load the configuration
	cfg, err := config.LoadConfig(configFilePath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Load the localization
	loc, err := internalui.LoadLocalization(cfg.Language)
	if err != nil {
		log.Fatalf("Error loading localization: %v", err)
	}

	// Start the UI
	err = ui.Main(func() {
		internalui.SetupUI(cfg, loc)
	})
	if err != nil {
		log.Fatalf("Error starting UI: %v", err)
	}
}
