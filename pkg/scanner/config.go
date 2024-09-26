package scanner

import (
	"encoding/json"
	"os"
)

type Config struct {
	Patterns  []string `json:"patterns"`
	FileTypes []string `json:"file_types"`
}

func LoadConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil

}
