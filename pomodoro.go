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

	formatter := formatSeconds
	switch {
	case wait >= 24*time.Hour:
		formatter = formatDays
	case wait >= time.Hour:
		formatter = formatHours
	case wait >= time.Minute:
		formatter = formatMinutes
	}

	fmt.Printf("Start timer for %s.\n\n", wait)

	if *simple {
		simpleCountdown(finish, formatter)
	} else {
		fullscreenCountdown(start, finish, formatter)
	}

	if !*silence {
		fmt.Println("\a") // \a is the bell literal.
	} else {
		fmt.Println()
	}
}
