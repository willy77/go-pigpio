package GoPiGPIO

import (
	"fmt"
)

type ScriptBuilder struct {
	code string
}

func (sc *ScriptBuilder) Code() string { return sc.code }

func (sc *ScriptBuilder) append(cmd string) *ScriptBuilder {
	sc.code += cmd + " "
	return sc
}

// Add
//
// Add x to accumulator
// A+=x; F=A
func (sc *ScriptBuilder) Add(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("add %d", x)) }

// And
//
// And x with accumulator
// A&=x; F=A
func (sc *ScriptBuilder) And(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("and %d", x)) }

// Call
//
// Call subroutine at tag L
// push(PC+1); PC=L
func (sc *ScriptBuilder) Call(tag int) *ScriptBuilder { return sc.append(fmt.Sprintf("call %d", tag)) }

// Cmp
//
// Compare x with accumulator
// F=A-x
func (sc *ScriptBuilder) Cmp(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("cmp %d", x)) }

// Dcr
//
// Decrement register
// --*y; F=*y
func (sc *ScriptBuilder) Dcr(y int) *ScriptBuilder { return sc.append(fmt.Sprintf("dcr %d", y)) }

// Dcra
//
// Decrement accumulator
// --A; F=A
func (sc *ScriptBuilder) Dcra() *ScriptBuilder { return sc.append("dcra") }

// Div
//
// Divide x into accumulator
// A/=x; F=A
func (sc *ScriptBuilder) Div(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("div %d", x)) }

// Evtwt
//
// Wait for an event to occur
// A=wait(x); F=A
func (sc *ScriptBuilder) Evtwt() *ScriptBuilder { return sc.append("evtwt") }

// Halt
//
// Halt
func (sc *ScriptBuilder) Halt() *ScriptBuilder { return sc.append("halt") }

// Inr
//
// Increment register
// ++*y; F=*y
func (sc *ScriptBuilder) Inr(y int) *ScriptBuilder { return sc.append(fmt.Sprintf("inr %d", y)) }

// Inra
//
// Increment accumulator
// ++A; F=A
func (sc *ScriptBuilder) Inra() *ScriptBuilder { return sc.append("inra") }

// Jm
//
// Jump if minus to tag L
// if (F<0) PC=L
func (sc *ScriptBuilder) Jm(tag int) *ScriptBuilder { return sc.append(fmt.Sprintf("jm %d", tag)) }

// Jmp
//
// Jump to tag L
// PC=L
func (sc *ScriptBuilder) Jmp(tag int) *ScriptBuilder { return sc.append(fmt.Sprintf("jmp %d", tag)) }

// Jnz
//
// Jump if non-zero to tag L
// if (F) PC=L
func (sc *ScriptBuilder) Jnz(tag int) *ScriptBuilder { return sc.append(fmt.Sprintf("jnz %d", tag)) }

// Jp
//
// Jump if positive to tag L
// if (F>=0) PC=L
func (sc *ScriptBuilder) Jp(tag int) *ScriptBuilder { return sc.append(fmt.Sprintf("jp %d", tag)) }

// Jz
//
// Jump if zero to tag L
// if (!F) PC=L
func (sc *ScriptBuilder) Jz(tag int) *ScriptBuilder { return sc.append(fmt.Sprintf("jz %d", tag)) }

// Ld
//
// Load register with x
// *y=x
func (sc *ScriptBuilder) Ld(y int, x int) *ScriptBuilder {
	return sc.append(fmt.Sprintf("ld %d %d", y, x))
}

// Lda
//
// Load accumulator with x
// A=x
func (sc *ScriptBuilder) Lda(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("lda %d", x)) }

// Mlt
//
// Multiply x with accumulator
// A*=x; F=A
func (sc *ScriptBuilder) Mlt(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("mlt %d", x)) }

// Mod
//
// Modulus x with accumulator
// A%=x; F=A
func (sc *ScriptBuilder) Mod(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("mod %d", x)) }

// Or
//
// Or x with accumulator
// A|=x; F=A
func (sc *ScriptBuilder) Or(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("or %d", x)) }

// Pop
//
// Pop register
// y=pop()
func (sc *ScriptBuilder) Pop(y int) *ScriptBuilder { return sc.append(fmt.Sprintf("pop %d", y)) }

// Popa
//
// Pop accumulator
// A=pop()
func (sc *ScriptBuilder) Popa() *ScriptBuilder { return sc.append("popa") }

// Push
//
// Push register
// push(y)
func (sc *ScriptBuilder) Push(y int) *ScriptBuilder { return sc.append(fmt.Sprintf("push %d", y)) }

// Pusha
//
// Push accumulator
// push(A)
func (sc *ScriptBuilder) Pusha() *ScriptBuilder { return sc.append("pusha") }

// Ret
//
// Return from subroutine
// PC=pop()
func (sc *ScriptBuilder) Ret() *ScriptBuilder { return sc.append("ret") }

// Rl
//
// Rotate left register x bits
// *y<<=x; F=*y
func (sc *ScriptBuilder) Rl(y int, x int) *ScriptBuilder {
	return sc.append(fmt.Sprintf("rl %d %d", y, x))
}

// Rla
//
// Rotate left accumulator x bits
// A<<=x; F=A
func (sc *ScriptBuilder) Rla(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("rla %d", x)) }

// Rr
//
// Rotate right register x bits
// *y>>=x; F=*y
func (sc *ScriptBuilder) Rr(y int, x int) *ScriptBuilder {
	return sc.append(fmt.Sprintf("rr %d %d", y, x))
}

// Rra
//
// Rotate right accumulator x bits
// A>>=x; F=A
func (sc *ScriptBuilder) Rra(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("rra %d", x)) }

// Sta
//
// Store accumulator in register
// y=A
func (sc *ScriptBuilder) Sta(y int) *ScriptBuilder { return sc.append(fmt.Sprintf("sta %d", y)) }

// Sub
//
// Subtract x from accumulator
// A-=x; F=A
func (sc *ScriptBuilder) Sub(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("sub %d", x)) }

// Sys
//
// Run external script (/opt/pigpio/cgi/str)
// system(str); F=A
func (sc *ScriptBuilder) Sys(str string) *ScriptBuilder { return sc.append(fmt.Sprintf("sys %s", str)) }

// Tag
//
// Label the current script position
// N/A
func (sc *ScriptBuilder) Tag(tag int) *ScriptBuilder { return sc.append(fmt.Sprintf("tag %d", tag)) }

// Wait
//
// Wait for a GPIO in x to change state
// A=wait(x); F=A
func (sc *ScriptBuilder) Wait(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("wait %d", x)) }

// Write
//
// Write level to gpio
func (sc *ScriptBuilder) Write(gpio *GpioPin, level GpioLevel) *ScriptBuilder {
	return sc.append(fmt.Sprintf("w %d %d", gpio.pin, level))
}

// X
//
// Exchange contents of registers y1 and y2
// t=*y1;*y1=*y2;*y2=t
func (sc *ScriptBuilder) X(y1 int, y2 int) *ScriptBuilder {
	return sc.append(fmt.Sprintf("x %d %d", y1, y2))
}

// Xa
//
// Exchange contents of accumulator and register
// t=A;A=*y;*y=t
func (sc *ScriptBuilder) Xa(y int) *ScriptBuilder { return sc.append(fmt.Sprintf("xa %d", y)) }

// Xor
//
// Xor x with accumulator
// A^=x; F=A
func (sc *ScriptBuilder) Xor(x int) *ScriptBuilder { return sc.append(fmt.Sprintf("xor %d", x)) }
