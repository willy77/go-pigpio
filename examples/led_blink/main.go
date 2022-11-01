package main

import (
	"GoPiGPIO/pkg/pigpio"
	"time"
)

func main() {
	var e error
	pi, e := pigpio.Initialize("localhost", 8888)
	if e != nil {
		panic(e)
	}

	gpio := pi.Gpio(2)
	if e = gpio.SetMode(pigpio.ModeOutput); e != nil {
		panic(e)
	}

	if e = gpio.Write(1); e != nil {
		panic(e)
	}

	time.Sleep(5 * time.Second)

	if e = gpio.Write(0); e != nil {
		panic(e)
	}

	if e = pi.Close(); e != nil {
		panic(e)
	}
}
