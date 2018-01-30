package main

import (
    "os"
    "io"
    "bufio"
    "encoding/csv"
    "fmt"
    "log"
    "time"
)

func GetUpsTrackings(filename string) []Tracking {
	// open csv file
    csvFile, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
	defer csvFile.Close()

    reader := csv.NewReader(bufio.NewReader(csvFile))

	today := time.Now().Format("2016-01-02")

    var trackings []Tracking

    for {
        fields, err := reader.Read()

        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatalln(err)
        }

		datetime := fields[3] // "20170203143750"
		y := datetime[0:4]
		m := datetime[4:6]
		d := datetime[6:8]
		shipDate := fmt.Sprintf("%s-%s-%s", y, m, d)

		if shipDate != today {
			continue
		}

        trackings = append(trackings, Tracking{
            OrderId:     fields[1],
            TrackingNum: fields[2],
            ShipDate:    shipDate,
        })
    }

	return trackings
}
