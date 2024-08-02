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

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
