package main

import (
	"embed"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:embed sauce/*
var df embed.FS

const (
	delayKeyfetchMS = 5
)

func main() {
	debug := flag.Bool("debug", false, "enable output")
	flag.Parse()

	dir, _ := df.ReadDir("sauce")
	saucestreamers := make([]beep.StreamSeekCloser, 0)
	for _, file := range dir {
		f, _ := df.Open("sauce/" + file.Name())
		streamer, format, _ := mp3.Decode(f)
		defer streamer.Close()
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		saucestreamers = append(saucestreamers, streamer)
	}
	kl := Keylogger{}
	for {
		key := kl.GetKey()
		if !key.Empty {
			rngval := rand.Intn(len(saucestreamers))
			if *debug {
				fmt.Println(strconv.Itoa(key.Keycode) + "   #" + strconv.Itoa(rngval))
			}
			speaker.Play(saucestreamers[rngval])
			saucestreamers[rngval].Seek(0)
		}
		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}
