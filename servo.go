package GoPiGPIO

type GpioServo struct {
	gpio *GpioPin
}

func (s *GpioServo) Gpio() *GpioPin { return s.gpio }

func (s *GpioServo) SetPulseWidth(width int) error {
	r, e := s.gpio.pi.socket.SendCommand(cmdServoPWS, s.gpio.pin, width, nil)
	if e != nil || r < 0 {
		return NewPiError(r, e, "SetPulseWidth(pin: %d, width: %d)", s.gpio.pin, width)
	}

	return nil
}

func (s *GpioServo) GetPulseWidth() (int, error) {
	r, e := s.gpio.pi.socket.SendCommand(cmdServoPWG, s.gpio.pin, 0, nil)
	if e != nil || r < 0 {
		return -1, NewPiError(r, e, "GetPulseWidth(pin: %d)", s.gpio.pin)
	}

	return r, nil
}
