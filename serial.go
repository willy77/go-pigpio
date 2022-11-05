package pigpio

type Serial struct {
    pi     *Pi
    handle int
    tty    string
    baud   int
}

func (s *Serial) TTY() string { return s.tty }
func (s *Serial) Baud() int   { return s.baud }
func (s *Serial) Handle() int { return s.handle }

func (s *Serial) Close() error {
    r, e := s.pi.socket.SendCommand(cmdSerialClose, s.handle, 0, nil)
    if e != nil || r < 0 {
        return newPiError(r, e, "Close(handle: %d)", s.handle)
    }

    return nil
}

func (s *Serial) Read(count int) ([]byte, error) {
    r, e := s.pi.socket.SendCommand(cmdSerialRead, s.handle, count, nil)
    if e != nil || r < 0 {
        return nil, newPiError(r, e, "Read(handle: %d, count: %d)", s.handle, count)
    }

    if r > 0 {
        data, re := s.pi.socket.Read(r)
        if re != nil {
            return nil, newPiError(r, re, "Read(handle: %d, count: %d)", s.handle, count)
        }

        return data, nil
    }

    return []byte{}, nil
}

func (s *Serial) ReadByte() (byte, error) {
    r, e := s.pi.socket.SendCommand(cmdSerialReadByte, s.handle, 0, nil)
    if e != nil || r < 0 {
        return 0, newPiError(r, e, "ReadByte(handle: %d)", s.handle)
    }

    return byte(r), nil
}

func (s *Serial) Write(data []byte) error {
    r, e := s.pi.socket.SendCommand(cmdSerialWrite, s.handle, 0, data)
    if e != nil || r < 0 {
        return newPiError(r, e, "Write(handle: %d, data: %v)", s.handle, data)
    }

    return nil
}

func (s *Serial) WriteByte(data byte) error {
    r, e := s.pi.socket.SendCommand(cmdSerialWriteByte, s.handle, int(data), nil)
    if e != nil || r < 0 {
        return newPiError(r, e, "WriteByte(handle: %d, data: %d)", s.handle, data)
    }

    return nil
}

func (s *Serial) DataAvailable() (bool, error) {
    r, e := s.pi.socket.SendCommand(cmdSerialAvailable, s.handle, 0, nil)
    if e != nil || r < 0 {
        return false, newPiError(r, e, "DataAvailable(handle: %d)", s.handle)
    }

    return r > 0, nil
}
