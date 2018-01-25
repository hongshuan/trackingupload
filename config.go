package main

import (
    "os"
    "encoding/json"
	"io/ioutil"
)

type Config struct {
	Carrier  string `json:"Carrier"`
	Filename string `json:"Filename"`
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
