// pormodoro timer
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func init() {
	const usage = `Usage of pomodoro:

	pomodoro [duration]

Duration defaults to 15 minutes. Durations may be expressed as integer minutes
(e.g. "15") or time with units (e.g. "1m30s" or "90s").
`
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}
	flag.Parse()
}

func getWaitDuration() time.Duration {
	// Default time of 15 minutes
	const defaultDuration = 15 * time.Minute

	if len(flag.Args()) > 1 {
		flag.Usage()
		os.Exit(2)
	}

	if flag.Arg(0) == "" {
		return defaultDuration
	}

	if n, err := strconv.Atoi(flag.Arg(0)); err == nil {
		return time.Duration(n) * time.Minute
	}

	if d, err := time.ParseDuration(flag.Arg(0)); err == nil {
		return d
	}

	flag.Usage()
	os.Exit(2)
	panic("missing return at end of function")
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
