// pormodoro timer
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	wait, err := waitDuration()
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}

	finish := start.Add(wait)

	switch {
	case wait >= 24*time.Hour:
		formatter = formatDays
	case wait >= time.Hour:
		formatter = formatHours
	case wait >= time.Minute:
		formatter = formatMinutes
	default:
		formatter = formatSeconds
	}

	fmt.Printf("Start timer for %s.\n\n", wait)

	if *simple {
		simpleCountdown(finish)
	} else {
		fullscreenCountdown(start, finish)
	}

	if !*silence {
		fmt.Println("\a") // \a is the bell literal.
	} else {
		fmt.Println()
	}
}
