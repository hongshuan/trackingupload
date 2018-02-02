package main

import (
    "os"
    "encoding/json"
	"io/ioutil"
)

type Carrier struct {
	Name        string `json:"Name"`
	Filename    string `json:"Filename"`
	Addressbook string `json:"Addressbook"`
}

type Config struct {
    Server         string    `json:"Server"`
    TrackingUrl    string    `json:"TrackingUrl"`
    AddressbookUrl string    `json:"AddressbookUrl"`
	Carriers       []Carrier `json:"Carriers"`
}

func LoadConfig(filename string) Config {
	input, err := os.Open(filename);
	CheckError(err)

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	CheckError(err)

	var cfg Config
	err = json.Unmarshal(jsonBytes, &cfg)
	CheckError(err)

	return cfg
}

func (config *Config) GetTrackingUrl() string {
	if len(config.Server) == 0 {
		config.Server = "http://localhost"
	}

	if len(config.TrackingUrl) == 0 {
		config.TrackingUrl = "/data/addressbook"
	}

	return config.Server + config.TrackingUrl
}

func (config *Config) GetAddressbookUrl() string {
	if len(config.Server) == 0 {
		config.Server = "http://localhost"
	}

	if len(config.AddressbookUrl) == 0 {
		config.AddressbookUrl = "/shipment/tracking"
	}

	return config.Server + config.AddressbookUrl
}
