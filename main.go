package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func startTicker(bpm, beats int, done chan struct{}) {
	streamer, format := loadAsset()
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/100))

	ticker := time.NewTicker(time.Duration(60000/bpm) * time.Millisecond)
	tickCounter := 0
	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			volume := 1.0
			if tickCounter%beats == 0 {
				volume = 2.0
			}
			vol := &effects.Volume{
				Streamer: streamer,
				Base:     2,
				Volume:   volume,
				Silent:   false,
			}
			speaker.Play(vol)

			streamer, _ = loadAsset()
			fmt.Printf("tick at %v\n", t)
			tickCounter++
		}
	}
}

//go:embed click.mp3
var data []byte

func loadAsset() (beep.StreamSeekCloser, beep.Format) {
	streamer, format, err := mp3.Decode(io.NopCloser(bytes.NewReader(data)))
	if err != nil {
		log.Fatal(err)
	}
	return streamer, format
}

func main() {
	bpmPtr := flag.Int("bpm", 60, "Beats per minute")
	beatsPtr := flag.Int("beats", 4, "Beats per bar")
	flag.Parse()

	done := make(chan struct{})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			done <- struct{}{}
		}
	}()
	startTicker(*bpmPtr, *beatsPtr, done)
}
