package ui

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed locales/en.json
var enJSON []byte

//go:embed locales/ru.json
var ruJSON []byte

type Localization struct {
	translations map[string]string
}

func LoadLocalization(language string) (*Localization, error) {
	var data []byte

	switch language {
	case "en":
		data = enJSON
	case "ru":
		data = ruJSON
	default:
		return nil, fmt.Errorf("unsupported language: %s", language)
	}

	var translations map[string]string
	err := json.Unmarshal(data, &translations)
	if err != nil {
		return nil, err
	}

	return &Localization{translations: translations}, nil
}

func (l *Localization) Translate(key string) string {
	return l.translations[key]
}
