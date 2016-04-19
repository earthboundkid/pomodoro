package main

import (
	"fmt"
	"os"
	"time"
)

func simpleCountdown(target time.Time) {
	for range time.Tick(100 * time.Millisecond) {
		timeLeft := -time.Since(target)
		if timeLeft < 0 {
			fmt.Print("Countdown: ", formatter(0), "   \r")
			return
		}
		fmt.Fprint(os.Stdout, "Countdown: ", formatter(timeLeft), "   \r")
		os.Stdout.Sync()
	}
}
