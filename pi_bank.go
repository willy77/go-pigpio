package pigpio

type PiBank byte

const (
	Bank1 PiBank = 1
	Bank2 PiBank = 2
)

func (p *Pi) ReadBank(bank PiBank) (int, error) {
	cmd := cmdBank1R
	if bank == Bank2 {
		cmd = cmdBank2R
	}

	r, e := p.socket.SendCommand(cmd, 0, 0, nil)
	if e != nil || r < 0 {
		return 0, newPiError(r, e, "ReadBank(bank: %d)", cmd)
	}

	return r, nil
}

func (p *Pi) ClearBank(bank PiBank, mask int) (int, error) {
	cmd := cmdBank1C
	if bank == Bank2 {
		cmd = cmdBank2C
	}

	r, e := p.socket.SendCommand(cmd, mask, 0, nil)
	if e != nil || r < 0 {
		return 0, newPiError(r, e, "ClearBank(bank: %d, mask: %d)", cmd, mask)
	}

	return r, nil
}

func (p *Pi) SetBank(bank PiBank, mask int) (int, error) {
	cmd := cmdBank1S
	if bank == Bank2 {
		cmd = cmdBank2S
	}

	r, e := p.socket.SendCommand(cmd, mask, 0, nil)
	if e != nil || r < 0 {
		return 0, newPiError(r, e, "SetBank(bank: %d, mask: %d)", cmd, mask)
	}

	return r, nil
}
