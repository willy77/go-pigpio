package GoPiGPIO

type WaveChain struct {
	pi   *Pi
	data []byte
}

func (wc *WaveChain) Bytes() []byte { return wc.data }

func (wc *WaveChain) append(data []byte) *WaveChain {
	wc.data = append(wc.data, data[:]...)
	return wc
}

func (wc *WaveChain) Wave(w *Wave) *WaveChain {
	return wc.append(convertToBytes(w.handle))
}

func (wc *WaveChain) LoopBegin() *WaveChain {
	return wc.append([]byte{255, 0})
}

// TODO wrong loop and delay (x + y*256)
func (wc *WaveChain) LoopEnd(repeat int) *WaveChain {
	return wc.append([]byte{255, 0, byte(repeat), byte(repeat << 8)})
}

func (wc *WaveChain) Delay(us int) *WaveChain {
	return wc.append([]byte{255, 0, byte(us), byte(us << 8)})
}

func (wc *WaveChain) LoopForever() *WaveChain {
	return wc.append([]byte{255, 3})
}

func (wc *WaveChain) Run() error {
	r, e := wc.pi.socket.SendCommand(cmdWaveChain, 0, 0, wc.Bytes())
	if e != nil || r < 0 {
		return NewPiError(r, e, "WaveChain.Run()")
	}

	return nil
}
