package main

import (
	"GoPiGPIO/pkg/pigpio"
	"fmt"
	"sync"
)

func main() {
	var e error
	pi, e := pigpio.Initialize("192.168.178.218", 8888)
	if e != nil {
		panic(e)
	}

	pi.Gpio(17).AddCallback(pigpio.EdgeRising, GpioCallback)

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

	if e = pi.Close(); e != nil {
		panic(e)
	}
}

func GpioCallback(gpio *pigpio.GpioPin, tick uint) {
	fmt.Println("GpioPin Callback")
}
