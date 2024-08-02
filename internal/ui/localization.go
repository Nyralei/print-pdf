package ui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Localization struct {
	translations map[string]string
}

func LoadLocalization(language string) (*Localization, error) {
	filePath := fmt.Sprintf("locales/%s.json", language)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var translations map[string]string
	err = json.Unmarshal(bytes, &translations)
	if err != nil {
		return nil, err
	}

	return &Localization{translations: translations}, nil
}

func (l *Localization) Translate(key string) string {
	return l.translations[key]
}
