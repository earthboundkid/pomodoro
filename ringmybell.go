package main

import (
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"

	"log"
	"os"
	"time"
)

const bell = "assets/ring_my_bell.mp3"

func ringMyBell() {

	f, err := os.Open(bell)
	if err != nil {
		log.Fatalf("could not play the file %s: %s", bell, err)
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		log.Fatalf("could not get mp3 decoder: %s", err)
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		log.Fatalf("could not get oto context: %s", err)
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

}
