package GoPiGPIO

type Wave struct {
	pi     *Pi
	handle int
}

func (w *Wave) Handle() int { return w.handle }

func (w *Wave) Delete() error {
	r, e := w.pi.socket.SendCommand(cmdWaveDelete, w.handle, 0, nil)
	if e != nil || r < 0 {
		return NewPiError(r, e, "Delete(handle: %d)", w.handle)
	}

	return nil
}

func (w *Wave) Once() (int, error) {
	r, e := w.pi.socket.SendCommand(cmdWaveOnce, w.handle, 0, nil)
	if e != nil || r < 0 {
		return -1, NewPiError(r, e, "SendOnce(handle: %d)", w.handle)
	}

	return r, nil
}

func (w *Wave) Repeat() (int, error) {
	r, e := w.pi.socket.SendCommand(cmdWaveRepeat, w.handle, 0, nil)
	if e != nil || r < 0 {
		return -1, NewPiError(r, e, "SendRepeat(handle: %d)", w.handle)
	}

	return r, nil
}

func (w *Wave) IsBusy() bool {
	r, e := w.pi.socket.SendCommand(cmdWaveBusy, w.handle, 0, nil)
	if e != nil || r < 0 {
		return false
	}

	return r == 1
}

func (w *Wave) Stop() error {
	r, e := w.pi.socket.SendCommand(cmdWaveHalt, 0, 0, nil)
	if e != nil || r < 0 {
		return NewPiError(r, e, "Stop(handle: %d)", w.handle)
	}

	return nil
}
