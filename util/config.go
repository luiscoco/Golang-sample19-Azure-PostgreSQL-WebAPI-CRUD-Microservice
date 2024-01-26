package util

import (
    "encoding/json"
    "os"
)

type Config struct {
    DatabaseURL string `json:"DatabaseURL"`
}

func LoadConfig(path string) (*Config, error) {
    configFile, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer configFile.Close()

    var config Config
    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}
