package main

import "fmt"
import "time"

func main() {

	quitChan := make(chan bool)

	// stop the timer after 5 seconds
	go func() {
		<-time.After(5 * time.Second)
		close(quitChan)
	}()

	// fire once per second
	t := time.NewTicker(time.Second)
	func() {
		for {
			select {
			case <-t.C:
			case <-quitChan:
				t.Stop()
				return
			}
			fmt.Println("Timer fired!")
		}
	}()
	fmt.Println("DONE")
}
