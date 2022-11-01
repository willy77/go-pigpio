package pigpio

func (p *Pi) OpenSPI(channel int, baud int, flags int) (*SPI, error) {
    r, e := p.socket.SendCommand(cmdSPIOpen, channel, baud, convertToBytes(flags))
    if e != nil || r < 0 {
        return nil, NewPiError(r, e, "OpenSPI(channel %d, baud: %d, flags: %d)", channel, baud, flags)
    }

    return &SPI{pi: p, handle: r, baud: baud, flags: flags, channel: channel}, nil
}
