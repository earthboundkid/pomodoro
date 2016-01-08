// pormodoro timer
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const defaultDuration = 25 * time.Minute

var silence = flag.Bool("silence", false, "")

func init() {
	const usage = `Usage of pomodoro:

	pomodoro [-silence] [duration]

Duration defaults to %d minutes. Durations may be expressed as integer minutes
(e.g. "15") or time with units (e.g. "1m30s" or "90s").

Chimes system bell at the end of the timer, unless -silence is set.
`
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, int(defaultDuration/time.Minute))
	}
	flag.Parse()
}

func waitDuration() (time.Duration, error) {
	if len(flag.Args()) > 1 {
		return 0, errors.New("Too many args...")
	}

	if flag.Arg(0) == "" {
		return defaultDuration, nil
	}

	if n, err := strconv.Atoi(flag.Arg(0)); err == nil {
		return time.Duration(n) * time.Minute, nil
	}

	if d, err := time.ParseDuration(flag.Arg(0)); err == nil {
		return d, nil
	}

	return 0, errors.New("Couldn't parse duration...")
}

func main() {
	wait, err := waitDuration()
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}

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

	if !*silence {
		fmt.Println("\a") // \a is the bell literal.
	} else {
		fmt.Println()
	}
}
