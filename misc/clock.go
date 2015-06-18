package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	location, err := time.LoadLocation("Local")
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatal(err)
	}
	target := time.Date(2015, 6, 21, 12, 37, 53, 0, location)
	fmt.Printf("\t%s\n", target.Format(time.UnixDate))

	var previous time.Time
	for {
		now := time.Now()
		now = now.Add(time.Duration(-now.Nanosecond()))
		if now != previous {
			previous = now
			remaining := -target.Sub(now)
			sign := ""
			if remaining >= 0 {
				sign = "+"
			}
			fmt.Printf("\t%s  %s%v              \r",
				now.Format(time.UnixDate), sign, remaining)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
