package main

import (
    "bufio"
    "encoding/csv"
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io"
	"io/ioutil"
    "log"
    "os"
    "strings"
	"net/http"
	"net/url"
)

const APIURL = "http://localhost/shipment/tracking"

type Config struct {
	Carrier  string `json:"Carrier"`
	Filename string `json:"Filename"`
}

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
	config := loadConfig("TrackingUpload.cfg")
	carrierCode = config.Carrier

	fmt.Println("Carrier:", carrierCode)

	var trackings []Tracking

	switch(carrierCode) {
	case "UPS":
		trackings = getUpsTrackings(config.Filename)
	case "Fedex":
		trackings = getFedexTrackings(config.Filename)
	case "Canada Post":
		trackings = getCanadaPostTrackings(config.Filename)
	}

	uploadTrackings(trackings)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
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

func getUpsTrackings(filename string) []Tracking {
    csvFile, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
	defer csvFile.Close()

    reader := csv.NewReader(bufio.NewReader(csvFile))

    var trackings []Tracking
    for {
        fields, err := reader.Read()

        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatalln(err)
        }

		date := fields[3] // "20170203143750"
		y := date[0:4]
		m := date[4:6]
		d := date[6:8]

        trackings = append(trackings, Tracking{
            OrderId:     fields[1],
            TrackingNum: fields[2],
            ShipDate:    fmt.Sprintf("%s-%s-%s", y, m, d),
        })
    }

	return trackings
}

func getFedexTrackings(filename string) []Tracking {
    csvFile, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
	defer csvFile.Close()

    reader := csv.NewReader(bufio.NewReader(csvFile))

    var trackings []Tracking

	reader.Read() // skip title line

    for {
        fields, err := reader.Read()

        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatalln(err)
        }

		date := fields[2] // "01232018"
		y := date[4:8]
		m := date[0:2]
		d := date[2:4]

        trackings = append(trackings, Tracking{
            OrderId:     fields[len(fields)-4],
            TrackingNum: fields[1],
            ShipDate:    fmt.Sprintf("%s-%s-%s", y, m, d),
        })
    }

	return trackings
}

type DeliveryRequests struct {
	XMLName  xml.Name `xml:"delivery-requests"`
	Requests []DeliveryRequest `xml:"delivery-request"`
}

type DeliveryRequest struct {
	TrackingNumber string `xml:"delivery-spec>reference>item-id"`
	OrderNumber    string `xml:"delivery-spec>reference>customer-ref1"`
	MailingDate    string `xml:"settlement-details>mailing-date"`
}

func getCanadaPostTrackings(filename string) []Tracking {
	xmlFile, err := os.Open(filename)
	checkError(err)
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var delivery DeliveryRequests
	xml.Unmarshal(byteValue, &delivery)

    var trackings []Tracking
	for _, request := range delivery.Requests {
        trackings = append(trackings, Tracking{
            OrderId:     request.OrderNumber,
            TrackingNum: request.TrackingNumber,
            ShipDate:    request.MailingDate,
        })
	}

	return trackings
}
