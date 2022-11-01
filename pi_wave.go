package GoPiGPIO

func (p *Pi) CreateWave(padPercent int, pulses ...GenericPulse) (*Wave, error) {
	var r int
	var e error

	pulseData := make([]byte, 0)
	for _, pu := range pulses {
		pulseData = append(pulseData, pu.toBytes()[:]...)
	}

	r, e = p.socket.SendCommand(cmdWaveAddGeneric, 0, 0, pulseData)
	if e != nil || r < 0 {
		return nil, NewPiError(r, e, "CreateWave().GenericPulses")
	}

	if padPercent <= 0 {
		r, e = p.socket.SendCommand(cmdWaveCreate, 0, 0, nil)
	} else {
		r, e = p.socket.SendCommand(cmdWaveCAP, padPercent, 0, nil)
	}

	if e != nil || r < 0 {
		return nil, NewPiError(r, e, "CreateWave().WaveCreate")
	}

	w := Wave{pi: p, handle: r}
	return &w, nil
}

func (p *Pi) CreateSerialWave(padPercent int, pulses ...SerialPulse) (*Wave, error) {
	var r int
	var e error

	for _, sp := range pulses {
		r, e = p.socket.SendCommand(cmdWaveAddSerial, sp.Gpio.pin, sp.Baud,
			append(convertToBytes(sp.DataBits, sp.StopBits, sp.Offset), sp.Data[:]...))
		if e != nil || r < 0 {
			return nil, NewPiError(r, e, "CreateSerialWave().SerialPulses")
		}
	}

	if padPercent <= 0 {
		r, e = p.socket.SendCommand(cmdWaveCreate, 0, 0, nil)
	} else {
		r, e = p.socket.SendCommand(cmdWaveCAP, padPercent, 0, nil)
	}
	if e != nil || r < 0 {
		return nil, NewPiError(r, e, "CreateWave().WaveCreate")
	}

	w := Wave{pi: p, handle: r}
	return &w, nil
}

func (p *Pi) ClearWaves() error {
	r, e := p.socket.SendCommand(cmdWaveClear, 0, 0, nil)
	if e != nil || r < 0 {
		return NewPiError(r, e, "ClearWaves()")
	}

	return nil
}

func (p *Pi) MaxMicros() (int, error) {
	r, e := p.socket.SendCommand(cmdWaveMicros, 2, 0, nil)
	if e != nil || r < 0 {
		return -1, NewPiError(r, e, "MaxMicros()")
	}

	return r, nil
}

func (p *Pi) MaxPulses() (int, error) {
	r, e := p.socket.SendCommand(cmdWavePulse, 2, 0, nil)
	if e != nil || r < 0 {
		return -1, NewPiError(r, e, "MaxPulses()")
	}

	return r, nil
}

func (p *Pi) MaxCBS() (int, error) {
	r, e := p.socket.SendCommand(cmdWaveCBS, 2, 0, nil)
	if e != nil || r < 0 {
		return -1, NewPiError(r, e, "MaxCBS()")
	}

	return r, nil
}

func (p *Pi) CreateWaveChain() *WaveChain {
	return &WaveChain{pi: p, data: make([]byte, 0)}
}
