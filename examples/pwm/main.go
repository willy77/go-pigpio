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

	gpio := pi.Gpio(pigpio.Bcm04)
	if e = gpio.SetMode(pigpio.pigpio.ModeOutput); e != nil {
		panic(e)
	}
	for dc := 0; dc < 255; dc++ {
		if _, e = gpio.Pwm().SetDutyCycle(dc); e != nil {
			panic(e)
		}
		fmt.Println("DutyCycle: ", dc)
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)

	if e = gpio.Write(0); e != nil {
		panic(e)
	}

	if e = pi.Close(); e != nil {
		panic(e)
	}
}
