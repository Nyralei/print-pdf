package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Language      string `json:"language"`
	OpenInBrowser bool   `json:"open_in_browser"`
	Landscape     bool   `json:"landscape"`
	ImagePath     string `json:"image_path"`
	PDFPath       string `json:"pdf_path"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		Language:      "en",
		OpenInBrowser: true,
		Landscape:     false,
		ImagePath:     "",
		PDFPath:       "",
	}
}

// LoadConfig loads the configuration from the given path
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			// If config file does not exist, create it with default values
			defaultConfig := DefaultConfig()
			err = SaveConfig(path, defaultConfig)
			if err != nil {
				return nil, err
			}
			return defaultConfig, nil
		}
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		// If config is invalid, use default config
		return DefaultConfig(), nil
	}

	return config, nil
}

// SaveConfig saves the given configuration to the specified path
func SaveConfig(path string, config *Config) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(config)
}
