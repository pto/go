package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\t\x1b[1;36mJust Go!\x1b[0m")
	target := time.Date(2015, 7, 5, 0, 0, 0, 0, time.UTC)
	fmt.Printf("\t%s\n", target.Format(time.UnixDate))

	var previous time.Time
	var days int
	for {
		now := time.Now()
		now = now.Add(time.Duration(-now.Nanosecond())) // truncate to second
		if now != previous {
			previous = now
			remaining := target.Sub(now)
			sign := "-"
			if remaining < 0 {
				sign = "+" // now is after target
				remaining = -remaining
			}
			if remaining >= 24*time.Hour {
				days = int(remaining / (24 * time.Hour))
				remaining = remaining % (24 * time.Hour)
			}
			fmt.Printf("\t%s  %s", now.Format(time.UnixDate), sign)
			if days > 0 {
				fmt.Printf("%dd", days)
			}
			fmt.Printf("%v              \r", remaining)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
