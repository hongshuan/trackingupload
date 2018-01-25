package main

import (
    "os"
    "io"
    "bufio"
    "encoding/csv"
    "fmt"
    "log"
)

func GetUpsTrackings(filename string) []Tracking {
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
