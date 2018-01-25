package main

import (
    "fmt"
    "os"
    "io"
    "bufio"
    "encoding/csv"
    "log"
)

func GetFedexTrackings(filename string) []Tracking {
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
