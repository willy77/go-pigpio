package pigpio

import "C"

type I2C struct {
    pi      *Pi
    handle  int
    bus     int
    address int
    flags   I2CFlags
}

func (i *I2C) Bus() int        { return i.bus }
func (i *I2C) Address() int    { return i.address }
func (i *I2C) Flags() I2CFlags { return i.flags }
func (i *I2C) Handle() int     { return i.handle }

func (i *I2C) Close() error {
    r, e := i.pi.socket.SendCommand(cmdI2CClose, i.handle, 0, nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "I2C.Close(handle: %d, bus: %d, address: %d)", i.handle, i.bus, i.address)
    }

    return nil
}

func (i *I2C) ReadByte() (byte, error) {
    r, e := i.pi.socket.SendCommand(cmdI2CReadByte, i.handle, 0, nil)
    if e != nil || r < 0 {
        return 0, NewPiError(r, e, "I2C.ReadByte(handle: %d, bus: %d, address: %d)",
            i.handle, i.bus, i.address)
    }

    return byte(r), nil
}

func (i *I2C) WriteByte(val byte) error {
    r, e := i.pi.socket.SendCommand(cmdI2CWriteByte, i.handle, int(val), nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "I2C.WriteByte(handle: %d, bus: %d, address: %d, val: %d)",
            i.handle, i.bus, i.address, val)
    }

    return nil
}

func (i *I2C) WriteQuick(bit GpioLevel) error {
    r, e := i.pi.socket.SendCommand(cmdI2CWriteQuick, i.handle, int(bit), nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "I2C.WriteQuick(handle: %d, bus: %d, address: %d, bit: %d)",
            i.handle, i.bus, i.address, bit)
    }

    return nil
}

func (i *I2C) ReadRegisterByte(register int) (byte, error) {
    r, e := i.pi.socket.SendCommand(cmdI2CReadByteData, i.handle, register, nil)
    if e != nil || r < 0 {
        return 0, NewPiError(r, e, "I2C.RegisterReadByte(handle: %d, bus: %d, address: %d, register: %d)",
            i.handle, i.bus, i.address, register)
    }

    return byte(r), nil
}

func (i *I2C) WriteRegisterByte(register int, val byte) error {
    r, e := i.pi.socket.SendCommand(cmdI2CWriteByteData, i.handle, register, convertToBytes(int(val)))
    if e != nil || r < 0 {
        return NewPiError(r, e, "I2C.RegisterWriteByte(handle: %d, bus: %d, address: %d, register: %d, val: %d)",
            i.handle, i.bus, i.address, register, val)
    }

    return nil
}

func (i *I2C) ReadRegisterWord(register int) (int, error) {
    r, e := i.pi.socket.SendCommand(cmdI2CReadWordData, i.handle, register, nil)
    if e != nil || r < 0 {
        return 0, NewPiError(r, e, "I2C.RegisterReadWord(handle: %d, bus: %d, address: %d, register: %d)",
            i.handle, i.bus, i.address, register)
    }

    return r, nil
}

func (i *I2C) WriteRegisterWord(register int, val int16) error {
    r, e := i.pi.socket.SendCommand(cmdI2CWriteWordData, i.handle, register, convertToBytes(int(val)))
    if e != nil || r < 0 {
        return NewPiError(r, e, "I2C.RegisterWriteByte(handle: %d, bus: %d, address: %d, register: %d, val: %d)",
            i.handle, i.bus, i.address, register, val)
    }

    return nil
}

func (i *I2C) ReadRegisterBlock(register int) ([]byte, error) {
    var e error
    var r int
    var d []byte

    r, e = i.pi.socket.SendCommand(cmdI2CReadBlockData, i.handle, register, nil)
    if e != nil || r < 0 {
        return nil, NewPiError(r, e, "I2C.ReadRegisterBlock(handle: %d, bus: %d, address: %d, register: %d)",
            i.handle, i.bus, i.address, register)
    }

    if r > 0 {
        d, e = i.pi.socket.Read(r)
        if e != nil {
            return nil, NewPiError(0, e, "I2C.ReadRegisterBlock(handle: %d, bus: %d, address: %d, register: %d, len: %d)",
                i.handle, i.bus, i.address, register, r)
        }

        return d, nil
    }

    return []byte{}, nil
}

func (i *I2C) WriteRegisterBlock(register int, data []byte) error {
    if len(data) > 0 {
        r, e := i.pi.socket.SendCommand(cmdI2CReadBlockData, i.handle, register, data)
        if e != nil || r < 0 {
            return NewPiError(r, e, "I2C.WriteRegisterBlock(handle: %d, bus: %d, address: %d, register: %d, data: %v)",
                i.handle, i.bus, i.address, register, data)
        }
    }

    return nil
}

func (i *I2C) ReadDevice(count int) ([]byte, error) {
    var e error
    var r int
    var d []byte

    r, e = i.pi.socket.SendCommand(cmdI2CReadDevice, i.handle, count, nil)
    if e != nil || r < 0 {
        return nil, NewPiError(r, e, "I2C.ReadDevice(handle: %d, bus: %d, address: %d, count: %d)",
            i.handle, i.bus, i.address, count)
    }

    if r > 0 {
        d, e = i.pi.socket.Read(r)
        if e != nil {
            return nil, NewPiError(0, e, "I2C.ReadDevice(handle: %d, bus: %d, address: %d, count: %d, len: %d)",
                i.handle, i.bus, i.address, count, r)
        }

        return d, nil
    }

    return []byte{}, nil
}

func (i *I2C) WriteDevice(data []byte) error {
    if len(data) > 0 {
        r, e := i.pi.socket.SendCommand(cmdI2CWriteDevice, i.handle, 0, data)
        if e != nil || r < 0 {
            return NewPiError(r, e, "I2C.WriteDevice(handle: %d, bus: %d, address: %d, data: %v)",
                i.handle, i.bus, i.address, data)
        }
    }

    return nil
}

func (i *I2C) ProcessCall(register int, val int16) (int16, error) {
    r, e := i.pi.socket.SendCommand(cmdI2CProcessCall, i.handle, register, convertToBytes(int(val)))
    if e != nil || r < 0 {
        return 0, NewPiError(r, e, "I2C.RegisterWriteByte(handle: %d, bus: %d, address: %d, register: %d, val: %d)",
            i.handle, i.bus, i.address, register, val)
    }

    return int16(r), nil
}

func (i *I2C) BlockProcessCall(register int, data []byte) ([]byte, error) {
    var e error
    var r int
    var d []byte

    r, e = i.pi.socket.SendCommand(cmdI2CProcessCall, i.handle, register, data)
    if e != nil || r < 0 {
        return nil, NewPiError(r, e, "I2C.ReadDevice(handle: %d, bus: %d, address: %d, register: %d, data: %v)",
            i.handle, i.bus, i.address, register, data)
    }

    if r > 0 {
        d, e = i.pi.socket.Read(r)
        if e != nil {
            return nil, NewPiError(0, e, "I2C.ReadDevice(handle: %d, bus: %d, address: %d, register: %d, data: %v)",
                i.handle, i.bus, i.address, register, data)
        }

        return d, nil
    }

    return []byte{}, nil
}
