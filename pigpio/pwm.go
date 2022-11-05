package pigpio

type GpioPwm struct {
	gpio *GpioPin
}

func (pwm *GpioPwm) Gpio() *GpioPin { return pwm.gpio }

func (pwm *GpioPwm) SetDutyCycle(dutyCycle int) (int, error) {
	r, e := pwm.gpio.pi.socket.SendCommand(cmdPwmDutyCycleS, pwm.gpio.pin, dutyCycle, nil)
	if e != nil || r < 0 {
		return 0, newPiError(r, e, "SetDutyCylce(pin: %d, dutyCycle: %d)", pwm.gpio.pin, dutyCycle)
	}

	return r, nil
}

func (pwm *GpioPwm) GetDutyCycle() (int, error) {
	r, e := pwm.gpio.pi.socket.SendCommand(cmdPwmDutyCycleG, pwm.gpio.pin, 0, nil)
	if e != nil || r < 0 {
		return 0, newPiError(r, e, "GetDutyCylce(pin: %d)", pwm.gpio.pin)
	}

	return r, nil
}

func (pwm *GpioPwm) SetRange(max int) error {
	r, e := pwm.gpio.pi.socket.SendCommand(cmdPwmRangeS, pwm.gpio.pin, max, nil)
	if e != nil || r < 0 {
		return newPiError(r, e, "SetRange(pin: %d, max: %d)", pwm.gpio.pin, max)
	}

	return nil
}

func (pwm *GpioPwm) GetRange() (int, error) {
	r, e := pwm.gpio.pi.socket.SendCommand(cmdPwmRangeG, pwm.gpio.pin, 0, nil)
	if e != nil || r < 0 {
		return 0, newPiError(r, e, "GetRange(pin: %d)", pwm.gpio.pin)
	}

	return r, nil
}

func (pwm *GpioPwm) GetRealRange() (int, error) {
	r, e := pwm.gpio.pi.socket.SendCommand(cmdPwmRealRangeG, pwm.gpio.pin, 0, nil)
	if e != nil || r < 0 {
		return 0, newPiError(r, e, "GetRealRange(pin: %d)", pwm.gpio.pin)
	}

	return r, nil
}

func (pwm *GpioPwm) SetFrequency(freq int) error {
	r, e := pwm.gpio.pi.socket.SendCommand(cmdPwmFreqS, pwm.gpio.pin, freq, nil)
	if e != nil || r < 0 {
		return newPiError(r, e, "SetFrequency(pin: %d, freq: %d)", pwm.gpio.pin, freq)
	}

	return nil
}

func (pwm *GpioPwm) GetFrequency() (int, error) {
	r, e := pwm.gpio.pi.socket.SendCommand(cmdPwmFreqG, pwm.gpio.pin, 0, nil)
	if e != nil || r < 0 {
		return 0, newPiError(r, e, "GetFrequency(pin: %d)", pwm.gpio.pin)
	}

	return r, nil
}
