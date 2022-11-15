package main

import (
	"fmt"
	"github.com/BxNiom/go-pigpio"
)

func main() {
	compiler := pigpio.NewCompiler()
	e := compiler.LoadDefaultMacros()
	if e != nil {
		panic(e)
	}

	script := "ld v0 1234\n" +
		"dec v0\n" +
		"ldr v1 v2 v3 v4"

	fmt.Println("Original script:\n" + script)
	sc, _ := compiler.Compile(script)
	fmt.Println("\nCompiled:\n" + sc)
}
