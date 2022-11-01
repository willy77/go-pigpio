package GoPiGPIO

func (p *Pi) StoreScript(code string) (*Script, error) {
	r, e := p.socket.SendCommand(cmdScript, 0, 0, []byte(code))
	if e != nil || r < 0 {
		return nil, NewPiError(r, e, "StoreScript(code: %s)", code)
	}

	return &Script{pi: p, handle: r, code: code}, nil
}
