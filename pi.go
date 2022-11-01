package GoPiGPIO

import (
	"time"
)

type Pi struct {
	socket      *Socket
	cbm         *callbackManager
	hwReversion uint
	gpioPins    []*GpioPin
}

func (p *Pi) Socket() *Socket   { return p.socket }
func (p *Pi) HwReversion() uint { return p.hwReversion }

// Initialize
//
// Create a new connection to pigpio daemon
// If connection to remote host failed make sure pigpiod-remote
// remote interface is enabled
//
// host: host ip address e.g. localhost, 192.168.178.100
//
// port: daemon port
func Initialize(host string, port int) (*Pi, error) {
	var e error
	var s, cs *Socket
	var rev int

	s, e = NewSocket(host, port)
	if e != nil {
		return nil, NewPiError(6, e, "Initialize(%s, %d)", host, port)
	}

	cs, e = NewSocket(host, port)
	if e != nil {
		return nil, NewPiError(7, e, "Initialize(%s, %d)", host, port)
	}

	time.Sleep(100 * time.Millisecond)
	rev, e = s.SendCommand(cmdGetVersion, 0, 0, nil)
	if e != nil {
		return nil, NewPiError(rev, e, "Initialize(%s, %d)", host, port)
	}

	pi := &Pi{socket: s, cbm: nil, hwReversion: uint(rev), gpioPins: make([]*GpioPin, 53)}
	pi.cbm, e = initializeCallbackManager(pi, cs)

	if e != nil {
		return nil, NewPiError(rev, e, "Initialize(%s, %d)", host, port)
	}

	return pi, nil
}

// Close
//
// Close socket connection to pigpiod
func (p *Pi) Close() error {
	if p.socket != nil {
		return p.socket.Close()
	}

	return nil
}

func (p *Pi) Gpio(bcm BCM) *GpioPin {
	pin := int(bcm)
	if p.gpioPins[pin] == nil {
		p.gpioPins[pin] = newGpioPin(p, pin)
	}

	return p.gpioPins[pin]
}

func (p *Pi) CurrentTick() (uint, error) {
	t, err := p.socket.SendCommand(cmdGetTick, 0, 0, nil)
	if err != nil || t < 0 {
		return 0, NewPiError(t, err, "CurrentTick()")
	}

	return uint(t), nil
}
