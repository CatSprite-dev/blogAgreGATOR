package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("%s", err)
	}
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error opening json file: %s", err)
	}

	var config Config
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return Config{}, fmt.Errorf("unmarshal error: %s", err)
	}
	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting the path od the home directory: %s", err)
	}
	filePath := homeDir + "/" + configFileName
	return filePath, nil

}
