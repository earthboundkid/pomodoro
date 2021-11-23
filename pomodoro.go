// pormodoro timer
package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
	"time"

	"github.com/dbatbold/beep"
)

// A small composition, especially for Pomodoro, by Alexander F. Rødseth, CC0 licensed
const beepNotes = "A9HRDE DQ C5q3r6i DI qw3rt67io0 DS C6qr6i0[ DS C5r[acn"

func playNotes() error {
	const (
		filenameOrEmpty = ""
		volume          = 100
	)
	music := beep.NewMusic(filenameOrEmpty)
	reader := bufio.NewReader(strings.NewReader(beepNotes))
	go music.Play(reader, volume)
	music.Wait()
	beep.FlushSoundBuffer()
	return nil
}

func main() {
	start := time.Now()

	var hasAudio bool
	if err := beep.OpenSoundDevice("default"); err == nil { // success
		if err := beep.InitSoundDevice(); err == nil { // success
			hasAudio = true
			defer beep.CloseSoundDevice()
		}
	}

	finish, err := waitDuration(start)
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	wait := finish.Sub(start)

	formatter := formatSeconds
	switch {
	case wait >= 24*time.Hour:
		formatter = formatDays
	case wait >= time.Hour:
		formatter = formatHours
	case wait >= time.Minute:
		formatter = formatMinutes
	}

	if *simple {
		simpleCountdown(finish, formatter)
	} else {
		fullscreenCountdown(start, finish, formatter)
	}

	if *silence {
		return
	}

	if hasAudio {
		playNotes()
	} else {
		beep.SendBell()
	}
}
