package main

import (
    "fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"net/url"
)

func UploadTrackings(carrierCode string, trackings []Tracking) []string {

	URL := config.GetTrackingUrl()

	messages := make([]string, 0)
	messages = append(messages, carrierCode)

	shipVia := "BTE"
	carrierName := ""

	for _, tracking := range trackings {
		form := url.Values{}
		form.Add("orderId",        tracking.OrderId)
		form.Add("shipDate",       tracking.ShipDate)
		form.Add("carrierCode",    carrierCode)
		form.Add("carrierName",    carrierName)
		form.Add("shipMethod",     tracking.ShipMethod)
		form.Add("trackingNumber", tracking.TrackingNum)
		form.Add("shipVia",        shipVia)

		req, err := http.NewRequest("POST", URL, strings.NewReader(form.Encode()))
		CheckError(err)

		req.PostForm = form
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		client := http.Client{}
		rsp, err := client.Do(req)
		CheckError(err)
		defer rsp.Body.Close()

		body, err := ioutil.ReadAll(rsp.Body)
		CheckError(err)

		fmt.Println(tracking.OrderId, tracking.TrackingNum, tracking.ShipDate, string(body))

		messages = append(messages, fmt.Sprintf("%s %s %s",
			tracking.OrderId, tracking.TrackingNum, tracking.ShipDate))
	}

	return messages
}
