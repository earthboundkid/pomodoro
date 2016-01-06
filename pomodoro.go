// pormodoro timer
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func getWaitDuration() time.Duration {
	// Default time of 15 minutes
	const defaultDuration = 15 * time.Minute

	if len(os.Args) == 1 {
		return defaultDuration
	}

	if n, err := strconv.Atoi(os.Args[1]); err == nil {
		return time.Duration(n) * time.Minute
	}

	if d, err := time.ParseDuration(os.Args[1]); err == nil {
		return d
	}

	return defaultDuration
}

func main() {
	wait := getWaitDuration()
	fmt.Printf("Start timer for %s.\n\n", wait)
	doneCh := time.After(wait)
	doneT := time.Now().Add(wait)

	go func() {
		for range time.Tick(200 * time.Millisecond) {
			duration := -time.Since(doneT)
			fmt.Printf("Countdown: %02d:%02d    \r",
				int(duration.Minutes()), int(duration.Seconds())%60)
		}
	}()
	<-doneCh
	fmt.Println("\a") // \a is the bell literal.
}
