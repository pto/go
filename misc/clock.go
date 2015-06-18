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
	tick := time.Tick(100 * time.Millisecond)
	for range tick {
		fmt.Printf("\t%s\r", time.Now().Format(time.UnixDate))
	}
}
