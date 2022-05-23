package main

import (
	"embed"
	"fmt"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:embed sauce/001.mp3
var df embed.FS

const (
	delayKeyfetchMS = 5
)

func main() {
	f, _ := df.Open("sauce/001.mp3")
	streamer, format, _ := mp3.Decode(f)
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	kl := Keylogger{}
	for {
		key := kl.GetKey()
		if !key.Empty {
			fmt.Println(key.Rune)
			speaker.Play(streamer)
			streamer.Seek(0)
		}
		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}
