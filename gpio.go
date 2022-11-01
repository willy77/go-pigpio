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

// cmdGW
//
// Set gpio pin to level
func (gp *GpioPin) Write(level GpioLevel) error {
    r, e := gp.pi.socket.SendCommand(cmdWrite, gp.pin, int(level), nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "cmdGW(pin: %d, level: %d)", gp.pin, level)
    }

    return nil
}

// cmdGR
//
// Returns level of the gpio pin
func (gp *GpioPin) Read() (GpioLevel, error) {
    r, e := gp.pi.socket.SendCommand(cmdRead, gp.pin, 0, nil)
    if e != nil || r < 0 {
        return Off, NewPiError(r, e, "cmdGR(pin: %d)", gp.pin)
    }

    return GpioLevel(r), nil
}

// SetMode
//
// Set mode of gpio pin
func (gp *GpioPin) SetMode(mode GpioMode) error {
    r, e := gp.pi.socket.SendCommand(cmdModeS, gp.pin, int(mode), nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "SetMode(pin: %d, mode: %d)", gp.pin, mode)
    }

    return nil
}

// GetMode
//
// Returns the current mode of the gpio pin
func (gp *GpioPin) GetMode() (GpioMode, error) {
    r, e := gp.pi.socket.SendCommand(cmdModeG, gp.pin, 0, nil)
    if e != nil || r < 0 {
        return -1, NewPiError(r, e, "GetMode(pin: %d)", gp.pin)
    }

    return GpioMode(r), nil
}

// SetPullMode
//
// Sets or clears the internal GPIO pull-up/down resistor.
func (gp *GpioPin) SetPullMode(mode PullUpDownMode) error {
    r, e := gp.pi.socket.SendCommand(cmdPUD, gp.pin, int(mode), nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "SetPullMode(pin: %d, mode: %d)", gp.pin, mode)
    }

    return nil
}

// Trigger
//
// Send a trigger pulse to a GPIO.  The GPIO is set to
// level for pulseLen (0-100) microseconds and then reset to not level.
func (gp *GpioPin) Trigger(pulseLen int, level GpioLevel) error {
    r, e := gp.pi.socket.SendCommand(cmdTrigger, gp.pin, pulseLen, convertToBytes(int(level)))
    if e != nil || r < 0 {
        return NewPiError(r, e, "Trigger(pin: %d, pulseLen: %d, level: %d)", gp.pin, pulseLen, level)
    }

    return nil
}

func (gp *GpioPin) SetNoiseFilter(steady int, active int) error {
    r, e := gp.pi.socket.SendCommand(cmdFilterNoise, steady, active, nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "SetNoiseFilter(gpio: %d, steady: %d, active: %d)", gp.pin, steady, active)
    }

    return nil
}

func (gp *GpioPin) SetGlitchFilter(steady int) error {
    r, e := gp.pi.socket.SendCommand(cmdFilterNoise, steady, 0, nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "SetGlitchFilter(gpio: %d, steady: %d)", gp.pin, steady)
    }

    return nil
}

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
