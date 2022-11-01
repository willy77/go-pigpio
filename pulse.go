package pigpio

type GenericPulse struct {
    GpioOn  *GpioPin
    GpioOff *GpioPin
    Delay   int
}

func CreatePulse(gpioOn *GpioPin, gpioOff *GpioPin, delay int) GenericPulse {
    return GenericPulse{GpioOn: gpioOn, GpioOff: gpioOff, Delay: delay}
}

func (p GenericPulse) toBytes() []byte {
    values := make([]int, 3)
    if p.GpioOn != nil {
        values[0] = 1 << p.GpioOn.pin
    } else {
        values[0] = 0
    }

    if p.GpioOff != nil {
        values[1] = 1 << p.GpioOff.pin
    } else {
        values[1] = 0
    }

    values[2] = p.Delay
    return convertToBytes(values...)
}

type SerialPulse struct {
    Gpio     *GpioPin
    Baud     int
    DataBits int
    StopBits int
    Offset   int
    Data     []byte
}

func CreateSerialPulse(gpio *GpioPin, baud int, data []byte) SerialPulse {
    return SerialPulse{Gpio: gpio, Baud: baud, Data: data, DataBits: 8, StopBits: 2, Offset: 0}
}

func CreateSerialPulseEx(gpio *GpioPin, baud int, data []byte, offset int, dataBits int, stopBits int) SerialPulse {
    return SerialPulse{Gpio: gpio, Baud: baud, Data: data, DataBits: dataBits, StopBits: stopBits, Offset: offset}
}
