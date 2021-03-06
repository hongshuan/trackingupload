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

var config Config

func handleRequest(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
		renderPage(w, r, nil)
    }

    if r.Method == http.MethodPost {
		r.ParseForm()
		if (r.FormValue("btn") == "Upload") {
			handleUpload(w, r)
		}
		if (r.FormValue("btn") == "Download") {
			handleDownload(w, r)
		}
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	messages := make([]string, 0)

	for _, carrier := range config.Carriers {
		carrierCode := carrier.Name

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

		msgs := UploadTrackings(carrierCode, trackings)

		messages = append(messages, msgs...)
		messages = append(messages, "\n")
	}

	renderPage(w, r, messages)
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	URL := config.GetAddressbookUrl()

	messages := make([]string, 0)

	for _, carrier := range config.Carriers {
		if len(carrier.Addressbook) == 0 {
			continue
		}

		var err error

		switch(carrier.Name) {
		case "UPS":
			err = DownloadFile(URL + "/ups", carrier.Addressbook)

		case "Fedex":
			err = DownloadFile(URL + "/fedex", carrier.Addressbook)

		case "Canada Post":
			err = DownloadFile(URL + "/canadapost", carrier.Addressbook)

		case "DHL":
			err = DownloadFile(URL + "/dhl", carrier.Addressbook)
		}

		if err == nil {
			messages = append(messages, "Addressbook downloaded for " + carrier.Name)
		}
	}

	renderPage(w, r, messages)
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
	config = LoadConfig("config.json")

    http.HandleFunc("/", handleRequest)

    fmt.Println("Listening on http://localhost:9090")

    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal(err)
    }
}
