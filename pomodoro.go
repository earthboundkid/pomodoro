// pormodoro timer
package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
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

func countdown(target time.Time, formatDuration func(time.Duration) string) {
	for range time.Tick(100 * time.Millisecond) {
		fmt.Print(formatDuration(-time.Since(target)))
	}
}

func formatDays(timeLeft time.Duration) string {
	days := int(timeLeft.Hours() / 24)
	hours := int(timeLeft.Hours()) % 24
	minutes := int(timeLeft.Minutes()) % 60
	seconds := int(timeLeft.Seconds()) % 60
	return fmt.Sprintf("Countdown: %d:%02d:%02d:%02d \r",
		days, hours, minutes, seconds)
}

func formatHours(timeLeft time.Duration) string {
	hours := int(timeLeft.Hours())
	minutes := int(timeLeft.Minutes()) % 60
	seconds := int(timeLeft.Seconds()) % 60
	return fmt.Sprintf("Countdown: %d:%02d:%02d\r",
		hours, minutes, seconds)
}

func formatMinutes(timeLeft time.Duration) string {
	minutes := int(timeLeft.Minutes())
	seconds := int(timeLeft.Seconds()) % 60
	return fmt.Sprintf("Countdown: %d:%02d\r", minutes, seconds)
}

func formatSeconds(timeLeft time.Duration) string {
	return fmt.Sprintf("Countdown: %02.1f \r", math.Abs(timeLeft.Seconds()))
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

	switch {
	case wait >= 24*time.Hour:
		go countdown(doneT, formatDays)
	case wait >= time.Hour:
		go countdown(doneT, formatHours)
	case wait >= time.Minute:
		go countdown(doneT, formatMinutes)
	default:
		go countdown(doneT, formatSeconds)
	}

	<-doneCh

	if !*silence {
		fmt.Println("\a") // \a is the bell literal.
	} else {
		fmt.Println()
	}
}
