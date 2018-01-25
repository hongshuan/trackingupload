package main

import (
    "log"
	"strings"
	"io/ioutil"
	"net/http"
	"net/url"
)

const APIURL = "http://localhost/shipment/tracking"

func uploadTrackings(trackings []Tracking) {
	for _, tracking := range trackings {
		form := url.Values{}
		form.Add("orderId",        tracking.OrderId)
		form.Add("shipDate",       tracking.ShipDate)
		form.Add("carrierCode",    carrierCode)
		form.Add("carrierName",    carrierName)
		form.Add("shipMethod",     tracking.ShipMethod)
		form.Add("trackingNumber", tracking.TrackingNum)
		form.Add("shipVia",        shipVia)

		req, err := http.NewRequest("POST", APIURL, strings.NewReader(form.Encode()))
		checkError(err)

		req.PostForm = form
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		client := http.Client{}
		rsp, err := client.Do(req)
		checkError(err)
		defer rsp.Body.Close()

		body, err := ioutil.ReadAll(rsp.Body)
		checkError(err)

		log.Println(tracking.OrderId, tracking.TrackingNum, tracking.ShipDate, string(body))
	}
}
