package pigpio

type I2CFlags uint16

// I2C flags
const (
    I2CWrite      I2CFlags = 0x0000 /* write data */
    I2CRead       I2CFlags = 0x0001 /* read data */
    I2CTen        I2CFlags = 0x0010 /* ten bit chip address */
    I2CReceiveLen I2CFlags = 0x0400 /* length will be first received byte */
    I2CNoReadAck  I2CFlags = 0x0800 /* if I2C_FUNC_PROTOCOL_MANGLING */
    I2CIgnoreNak  I2CFlags = 0x1000 /* if I2C_FUNC_PROTOCOL_MANGLING */
    I2CRevDirAddr I2CFlags = 0x2000 /* if I2C_FUNC_PROTOCOL_MANGLING */
    I2CNoStart    I2CFlags = 0x4000 /* if I2C_FUNC_PROTOCOL_MANGLING */
)

func (p *Pi) OpenI2C(bus int, address int, flags I2CFlags) (*I2C, error) {
    r, e := p.socket.SendCommand(cmdI2COpen, bus, address, convertToBytes(int(flags)))
    if e != nil || r < 0 {
        return nil, NewPiError(r, e, "OpenI2C(bus: %d, address: %d, flags: %d)", bus, address, flags)
    }

    return &I2C{pi: p, handle: r, bus: bus, address: address, flags: flags}, nil
}
