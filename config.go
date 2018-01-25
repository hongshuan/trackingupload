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

func loadConfig(filename string) Config {
	input, err := os.Open(filename);
	checkError(err)

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	checkError(err)

	var cfg Config
	err = json.Unmarshal(jsonBytes, &cfg)
	checkError(err)

	return cfg
}
