package util

import (
	"gopkg.in/yaml.v3"
	"os"
)

func FetchYAML(fileName string, config interface{}) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, config)
}
