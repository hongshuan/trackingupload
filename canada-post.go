package main

import (
    "os"
    "time"
	"io/ioutil"
    "encoding/xml"
)

type DeliveryRequests struct {
	XMLName  xml.Name          `xml:"delivery-requests"`
	Requests []DeliveryRequest `xml:"delivery-request"`
}

type DeliveryRequest struct {
	TrackingNumber string `xml:"delivery-spec>reference>item-id"`
	OrderNumber    string `xml:"delivery-spec>reference>customer-ref1"`
	MailingDate    string `xml:"settlement-details>mailing-date"`
}

func GetCanadaPostTrackings(filename string) []Tracking {
    var trackings []Tracking

    if !FileExists(filename) {
        return trackings
    }

	xmlFile, err := os.Open(filename)
	CheckError(err)
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var delivery DeliveryRequests
	xml.Unmarshal(byteValue, &delivery)

	today := time.Now().Format("2006-01-02")

	for _, request := range delivery.Requests {
		if request.MailingDate != today {
			continue
		}

        trackings = append(trackings, Tracking{
            OrderId:     request.OrderNumber,
            TrackingNum: request.TrackingNumber,
            ShipDate:    request.MailingDate,
        })
	}

	return trackings
}
