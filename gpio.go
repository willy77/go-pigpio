package pigpio

type GpioPin struct {
	pi           *Pi
	pin          int
	pwm          *GpioPwm
	servo        *GpioServo
	callbackInit bool
	callbacks    []Callback
}

func (gp *GpioPin) Pwm() *GpioPwm     { return gp.pwm }
func (gp *GpioPin) Servo() *GpioServo { return gp.servo }

func newGpioPin(pi *Pi, pin int) *GpioPin {
	gpio := GpioPin{pin: pin,
		pi:           pi,
		callbackInit: false,
		callbacks:    make([]Callback, 0)}

	gpio.pwm = &GpioPwm{gpio: &gpio}
	gpio.servo = &GpioServo{gpio: &gpio}

	return &gpio
}

// Write sets GPIO level
//
// Example:
//
//	gpio := pi.Gpio(17)
//	gpio.SetMode(ModeInput)
//
//	gpio.Write(Low)
//	fmt.Println(gpio.Read)
//	// Output: 0
//
//	gpio.Write(High)
//	fmt.Println(gpio.Read)
//	// Output: 1
func (gp *GpioPin) Write(level GpioLevel) error {
	r, e := gp.pi.socket.SendCommand(cmdWrite, gp.pin, int(level), nil)
	if e != nil || r < 0 {
		return newPiError(r, e, "cmdGW(pin: %d, level: %d)", gp.pin, level)
	}

	return nil
}

// Read returns GPIO level
//
// Example:
//
//	pi.Gpio(23).SetMode(ModeInput)
//	pi.Gpio(23).SetPullMode(pudDown)
//	fmt.Println(pi.Gpio(23).Read())
//	// Output:    0
func (gp *GpioPin) Read() (GpioLevel, error) {
	r, e := gp.pi.socket.SendCommand(cmdRead, gp.pin, 0, nil)
	if e != nil || r < 0 {
		return Off, newPiError(r, e, "cmdGR(pin: %d)", gp.pin)
	}

	return GpioLevel(r), nil
}

// SetMode sets GPIO mode
//
// Parameters:
//
//	mode: GpioMode
//
// Example:
//
//	pi.Gpio(4).SetMode(4, ModeInput)   // GPIO 4 as input
//	pi.Gpio(4).SetMode(17, ModeOutput) // GPIO 17 as output
//	pi.Gpio(4).SetMode(24, ModeAlt2)   // GPIO 24 as ALT2
func (gp *GpioPin) SetMode(mode GpioMode) error {
	r, e := gp.pi.socket.SendCommand(cmdModeS, gp.pin, int(mode), nil)
	if e != nil || r < 0 {
		return newPiError(r, e, "SetMode(pin: %d, mode: %d)", gp.pin, mode)
	}

	return nil
}

// GetMode returns the GPIO mode
func (gp *GpioPin) GetMode() (GpioMode, error) {
	r, e := gp.pi.socket.SendCommand(cmdModeG, gp.pin, 0, nil)
	if e != nil || r < 0 {
		return -1, newPiError(r, e, "GetMode(pin: %d)", gp.pin)
	}

	return GpioMode(r), nil
}

// SetPullMode sets or clears the internal GPIO pull-up/down resistor
//
// Parameters:
//
//	mode: PullUpDownMode
//
// Example:
//
//	pi.Gpio(17).SetPullMode(PudOff)
//	pi.Gpio(23).SetPullMode(PudUp)
//	pi.Gpio(24).SetPullMode(PudDown)
func (gp *GpioPin) SetPullMode(mode PullUpDownMode) error {
	r, e := gp.pi.socket.SendCommand(cmdPUD, gp.pin, int(mode), nil)
	if e != nil || r < 0 {
		return newPiError(r, e, "SetPullMode(pin: %d, mode: %d)", gp.pin, mode)
	}

	return nil
}

// Trigger send a trigger pulse to a GPIO. The GPIO is set to
// level for pulseLen (1-100) microseconds and then reset to not level.
//
// Example:
//
//	pi.Gpio(17).Trigger(10, High)
func (gp *GpioPin) Trigger(pulseLen int, level GpioLevel) error {
	r, e := gp.pi.socket.SendCommand(cmdTrigger, gp.pin, pulseLen, convertToBytes(int(level)))
	if e != nil || r < 0 {
		return newPiError(r, e, "Trigger(pin: %d, pulseLen: %d, level: %d)", gp.pin, pulseLen, level)
	}

	return nil
}

// SetNoiseFilter sets a noise filter on a GPIO.
//
// Level changes on the GPIO are ignored until a level which has
// been stable for steady microseconds is detected. Level
// changes on the GPIO are then reported for active microseconds
// after which the process repeats.
//
// Parameters:
//
//	steady: 0 - 300000
//	active: 0 - 1000000
//
// Example:
//
//	pi.Gpio(8).SetNoiseFilter(1000, 5000)
func (gp *GpioPin) SetNoiseFilter(steady int, active int) error {
	r, e := gp.pi.socket.SendCommand(cmdFilterNoise, steady, active, nil)
	if e != nil || r < 0 {
		return newPiError(r, e, "SetNoiseFilter(gpio: %d, steady: %d, active: %d)", gp.pin, steady, active)
	}

	return nil
}

// SetGlitchFilter sets a glitch filter on a GPIO.
//
// Level changes on the GPIO are not reported unless the level
// has been stable for at least steady microseconds. The level
// is then reported.  Level changes of less than steady
// microseconds are ignored.
//
// Parameters:
//
//	steady: 0 - 300000
//
// Example:
//
//	pi.Gpio(8).SetGlitchFilter(100)
func (gp *GpioPin) SetGlitchFilter(steady int) error {
	r, e := gp.pi.socket.SendCommand(cmdFilterNoise, steady, 0, nil)
	if e != nil || r < 0 {
		return newPiError(r, e, "SetGlitchFilter(gpio: %d, steady: %d)", gp.pin, steady)
	}

	return nil
}

// AddCallback calls a user supplied function (fn) whenever the
// specified GPIO edge is detected.
// The user supplied function receives two parameters, the GPIO and the tick.
func (gp *GpioPin) AddCallback(edge Edge, fn CallbackFunc) Callback {
	cb := Callback{edge: edge, fn: fn, handle: gp.pi.cbm.nextHandle(), bit: 1 << gp.pin}
	gp.callbacks = append(gp.callbacks, cb)

	if !gp.callbackInit {
		_ = gp.pi.cbm.append(1 << gp.pin)
		gp.callbackInit = true
	}

	return cb
}

func (gp *GpioPin) RemoveCallback(cb Callback) {
	for idx, cba := range gp.callbacks {
		if cba.handle == cb.handle {
			gp.callbacks = append(gp.callbacks[:idx], gp.callbacks[idx+1:]...)
		}
	}

	if len(gp.callbacks) <= 0 && gp.callbackInit {
		_ = gp.pi.cbm.remove(1 << gp.pin)
		gp.callbackInit = false
	}
}

func (gp *GpioPin) invokeCallbacks(level int, tick uint) {
	for _, cb := range gp.callbacks {
		if int(cb.edge) == level {
			cb.fn(gp, tick)
		}
	}
}
