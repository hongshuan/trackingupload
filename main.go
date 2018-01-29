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

	for _, carrier := range config.Carriers {
		carrierCode = carrier.Name

		fmt.Println("Carrier:", carrierCode)

		var trackings []Tracking

		switch(carrierCode) {
		case "UPS":
			trackings = GetUpsTrackings(carrier.Filename)
		case "Fedex":
			trackings = GetFedexTrackings(carrier.Filename)
		case "Canada Post":
			trackings = GetCanadaPostTrackings(carrier.Filename)
		}

		UploadTrackings(trackings)
	}
}
