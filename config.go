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
	Carriers []Carrier `json:"Carriers"`
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
