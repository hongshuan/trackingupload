package main

import (
    "fmt"
)

type Tracking struct {
    OrderId     string
    TrackingNum string
    ShipDate    string
    ShipMethod  string
}

var carrierCode string = "UPS"
var carrierName string = ""
var shipVia     string = "BTE"

func main() {
	config := LoadConfig("TrackingUpload.cfg")
	carrierCode = config.Carrier

	fmt.Println("Carrier:", carrierCode)

	var trackings []Tracking

	switch(carrierCode) {
	case "UPS":
		trackings = GetUpsTrackings(config.Filename)
	case "Fedex":
		trackings = GetFedexTrackings(config.Filename)
	case "Canada Post":
		trackings = GetCanadaPostTrackings(config.Filename)
	}

	UploadTrackings(trackings)
}
