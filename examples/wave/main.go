package main

import (
	"fmt"
	"github.com/BxNiom/go-pigpio/pkg/pigpio"
	"time"
)

func main() {
	var e error
	pi, e := pigpio.Initialize("localhost", 8888)
	if e != nil {
		panic(e)
	}

	p2 := pi.Gpio(2)
	p3 := pi.Gpio(3)
	p4 := pi.Gpio(4)

	p2.SetMode(pigpio.ModeOutput)
	p3.SetMode(pigpio.ModeOutput)
	p4.SetMode(pigpio.ModeOutput)

	if e = pi.ClearWaves(); e != nil {
		panic(e)
	}

	wave, e := pi.CreateWave(0,
		pigpio.GenericPulse{GpioOn: p2, GpioOff: p4, Delay: 500000},
		pigpio.GenericPulse{GpioOn: p3, GpioOff: p2, Delay: 500000},
		pigpio.GenericPulse{GpioOn: p4, GpioOff: p3, Delay: 500000},
		pigpio.GenericPulse{GpioOn: p2, GpioOff: p4, Delay: 500000})

	if e != nil {
		panic(e)
	}

	c, e := wave.Repeat()
	if e != nil {
		panic(e)
	}

	fmt.Println("DMA Controls: ", c)
	time.Sleep(10 * time.Second)

	if e = wave.Stop(); e != nil {
		panic(e)
	}

	if e = wave.Delete(); e != nil {
		panic(e)
	}

	p2.Write(0)
	p3.Write(0)
	p4.Write(0)

	if e = pi.Close(); e != nil {
		panic(e)
	}
}
