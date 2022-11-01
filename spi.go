package pigpio

type SPI struct {
    pi      *Pi
    handle  int
    flags   int
    baud    int
    channel int
}

func (s *SPI) Close() error {
    r, e := s.pi.socket.SendCommand(cmdSPIClose, s.handle, 0, nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "Close(handle: %d)", s.handle)
    }

    return nil
}

func (s *SPI) Read(count int) ([]byte, error) {
    r, e := s.pi.socket.SendCommand(cmdSPIRead, s.handle, count, nil)
    if e != nil || r < 0 {
        return nil, NewPiError(r, e, "Read(handle: %d, count: %d)", s.handle, count)
    }

    data, er := s.pi.socket.Read(r)
    if er != nil {
        return nil, NewPiError(-1, er, "Read(handle: %d, count: %d, bytes: %d)", s.handle, count, r)
    }

    return data, nil
}

func (s *SPI) Write(data []byte) error {
    r, e := s.pi.socket.SendCommand(cmdSPIWrite, s.handle, 0, data)
    if e != nil || r < 0 {
        return NewPiError(r, e, "Write(handle: %d, data: %v)", s.handle, data)
    }

    return nil
}

func (s *SPI) Xfer(data []byte) ([]byte, error) {
    r, e := s.pi.socket.SendCommand(cmdSPIXfer, s.handle, 0, data)
    if e != nil || r < 0 {
        return nil, NewPiError(r, e, "Xfer(handle: %d, data: %v)", s.handle, data)
    }

    data, er := s.pi.socket.Read(r)
    if er != nil {
        return nil, NewPiError(-1, er, "Xfer(handle: %d, data: %v, bytes: %d)", s.handle, data, r)
    }

    return data, nil
}
