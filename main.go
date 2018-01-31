package main

import (
    "fmt"
    "net/http"
    "html/template"
    "log"
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

func handleRequest(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
		renderPage(w, r, nil)
    }

    if r.Method == "POST" {
		handleUpload(w, r)
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	data := trackingUpload()
	renderPage(w, r, data)
}

func trackingUpload() []string {
	config := LoadConfig("config.json")

	messages := make([]string, 0)

	for _, carrier := range config.Carriers {
		// save to global variable
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

		msgs := UploadTrackings(trackings)

		messages = append(messages, msgs...)
		messages = append(messages, "\n")
	}

	return messages
}

func renderPage(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl, err := template.New("tpl").Parse(pageTpl)
	if err != nil {
		http.Error(w, "500 Internal Error.", 500)
		fmt.Println(err)
		return
	}

	err = tpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
    http.HandleFunc("/", handleRequest)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal(err)
    }
}
