package main

import (
	"machine"

	"log"
	"time"

	"github.com/pkg/errors"
	"tinygo.org/x/drivers/ws2812"

	"github.com/242617/m5atom/pkg/scenarios"
	"github.com/242617/m5atom/pkg/screen"
)

const (
	FPS      = 24
	FPSDelay = time.Second / FPS
)

func init() { log.SetFlags(log.Ltime) }
func main() {
	var ledsPin machine.Pin = 27
	ledsPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	var btnPin machine.Pin = 39
	btnPin.Configure(machine.PinConfig{Mode: machine.PinInput})

	leds := ws2812.New(ledsPin)

	screen, err := screen.New(
		screen.WithDimensions(5, 5),
		screen.WithBrightness(.2),
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "create screen"))
	}

	go scenarios.Flags(screen)

	for {
		if err := leds.WriteColors(screen.Render()); err != nil {
			log.Fatal(err)
		}
		time.Sleep(FPSDelay)
	}
}
