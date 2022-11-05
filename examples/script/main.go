package main

import (
	"fmt"
	"github.com/BxNiom/go-pigpio/pkg/pigpio"
	"time"
)

func main() {
	var e error
	pi, e := pigpio.Initialize(localhost, 8888)
	if e != nil {
		panic(e)
	}

	p2 := pi.Gpio(2)
	p3 := pi.Gpio(3)
	p4 := pi.Gpio(4)

	p2.SetMode(pigpio.pigpio.ModeOutput)
	p3.SetMode(pigpio.ModeOutput)
	p4.SetMode(pigpio.ModeOutput)

	sb := pigpio.ScriptBuilder{}
	sb.Lda(-50).
		Tag(9999).
		Write(p2, 1).
		Write(p2, 0).
		Write(p3, 1).
		Write(p3, 0).
		Write(p4, 1).
		Write(p4, 0).
		Inra().
		Jnz(9999).
		Halt()

	fmt.Printf("Script code: %s\n", sb.Code())
	scr, e := pi.StoreScript(sb.Code())

	if e != nil {
		panic(e)
	}

	fmt.Printf("Script handle: %d\n", scr.Handle())

	e = scr.Run()
	if e != nil {
		panic(e)
	}

	time.Sleep(5 * time.Second)

	if e = scr.Delete(); e != nil {
		panic(e)
	}

	if e = pi.Close(); e != nil {
		panic(e)
	}
}
