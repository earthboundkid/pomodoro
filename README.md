# pomodoro
Command line [pomodoro timer](https://en.wikipedia.org/wiki/Pomodoro_Technique), implemented in Go

## Installation
First install [Go](http://golang.org) and set your `GOPATH` environmental variable to the directory you would like the project saved in. Then run `go get github.com/carlmjohnson/pomodoro`. The binary will be installed in `$GOPATH/bin`. If you don't want to keep the source, you can instead run `GOPATH=/tmp/srp go get github.com/carlmjohnson/pomodoro && cp /tmp/srp/bin/pomodoro .` to install the binary to your current working directory.

## Usage
Usage of pomodoro:

    pomodoro [-silence] [duration]

Duration defaults to %d minutes. Durations may be expressed as integer minutes
(e.g. "15") or time with units (e.g. "1m30s" or "90s").

Chimes system bell at the end of the timer, unless -silence is set.

## Screenshots
```bash
$ pomodoro 1s
Start timer for 1s.

Countdown: 0.0

$ pomodoro
Start timer for 25m0s.

Countdown: 24:43
```
