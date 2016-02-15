package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/0xAX/notificator"
)

var duration string
var silence bool

var notify *notificator.Notificator

func init() {
  const (
    defaultDuration = "25m"
    durationUsage = "Duration, e.g. 1m30s or 90s"
    silenceUsage = "Suppresses the notification after the timer ends"
  )
  
  flag.StringVar(&duration, "duration", defaultDuration, durationUsage)
  flag.StringVar(&duration, "d", defaultDuration, durationUsage+" (shorthand)")
  flag.BoolVar(&silence, "silence", false, silenceUsage)
  flag.BoolVar(&silence, "s", false, silenceUsage+" (shorthand)")
	flag.Parse()
  
  notify = notificator.New(notificator.Options{
	 DefaultIcon: "icon/default.png",
	 AppName:     "Pomodoro",
  })
}

func waitDuration() (time.Duration, error) {
	if n, err := strconv.Atoi(duration); err == nil {
		return time.Duration(n) * time.Minute, nil
	}

	if d, err := time.ParseDuration(duration); err == nil {
		return d, nil
	}

	return 0, errors.New("Couldn't parse duration...")
}

func countdown(target time.Time, formatDuration func(time.Duration) string) {
	for range time.Tick(100 * time.Millisecond) {
		timeLeft := -time.Since(target)
		if timeLeft < 0 {
			fmt.Print(formatDuration(0))
			return
		}
		fmt.Fprint(os.Stdout, formatDuration(timeLeft))
		os.Stdout.Sync()
	}
}

func formatMinutes(timeLeft time.Duration) string {
	minutes := int(timeLeft.Minutes())
	seconds := int(timeLeft.Seconds()) % 60
	return fmt.Sprintf("Countdown: %d:%02d \r", minutes, seconds)
}

func main() {
	wait, err := waitDuration()
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	doneT := time.Now().Add(wait)

	fmt.Printf("Start timer for %s.\n\n", wait)
	countdown(doneT, formatMinutes)

	if !silence {
		notify.Push("Done!", "You finished your pomodoro", "", notificator.UR_NORMAL)
	}
}
