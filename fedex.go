package main

import (
    "fmt"
    "os"
    "io"
    "bufio"
    "encoding/csv"
    "log"
    "time"
)

func GetFedexTrackings(filename string) []Tracking {
	// open csv file
    csvFile, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
	defer csvFile.Close()

    reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Read() // skip title line

	today := time.Now().Format("2016-01-02")

    var trackings []Tracking

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

		shipDate := fmt.Sprintf("%s-%s-%s", y, m, d)

		if shipDate != today {
			continue
		}

        trackings = append(trackings, Tracking{
            OrderId:     fields[len(fields)-4],
            TrackingNum: fields[1],
            ShipDate:    shipDate,
        })
    }

	return trackings
}
