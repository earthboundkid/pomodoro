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

var silence = flag.Bool("silence", false, "Don't ring bell after countdown")

var simple = flag.Bool("simple", false, "Display simple countdown")

func init() {
	const usage = `Usage of pomodoro:

    pomodoro [options] [duration]

Duration defaults to %d minutes. Durations may be expressed as integer minutes
(e.g. "15") or time with units (e.g. "1m30s" or "90s").

Chimes system bell at the end of the timer, unless -silence is set.
`
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, int(defaultDuration/time.Minute))
		flag.PrintDefaults()
	}
	flag.Parse()
}

func waitDuration() (time.Duration, error) {
	if flag.NArg() > 1 {
		return 0, errors.New("Too many args...")
	}

	arg := flag.Arg(0)

	if arg == "" {
		return defaultDuration, nil
	}

	if n, err := strconv.Atoi(arg); err == nil {
		return time.Duration(n) * time.Minute, nil
	}

	if d, err := time.ParseDuration(arg); err == nil {
		return d, nil
	}

	return 0, errors.New("Couldn't parse duration...")
}
