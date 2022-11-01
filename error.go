package GoPiGPIO

import "fmt"

type PiError struct {
	Code        int
	Message     string
	CodeMessage string
	InnerError  error
}

func NewPiError(code int, innerError error, msgFmt string, a ...any) error {
	return &PiError{
		Code:        code,
		Message:     fmt.Sprintf(msgFmt, a),
		CodeMessage: errorCodeMessage(code),
		InnerError:  innerError}
}

func (e *PiError) Error() string {
	msg := fmt.Sprintf("%s (ErrorCode: %d; %s)", e.Message, e.Code, e.CodeMessage)
	if e.InnerError != nil {
		msg += fmt.Sprintf(": %s", e.InnerError.Error())
	}

	return msg
}

func errorCodeMessage(errorCode int) string {

	switch errorCode {
	case 1:
		return "could not resolve ip address"
	case 2:
		return "could not connect"
	case 3:
		return "socket write error"
	case 4:
		return "socket read error"
	case 5:
		return "could not read response"
	case 6:
		return "initialize failed"
	case 7:
		return "initialize callback socket failed"
	case 8:
		return "event already registered"
	case -1:
		return "gpioInitialise failed"
	case -2:
		return "Gpionot 0-31"
	case -3:
		return "Gpionot 0-53"
	case -4:
		return "mode not 0-7"
	case -5:
		return "level not 0-1"
	case -6:
		return "pud not 0-2"
	case -7:
		return "pulsewidth not 0 or 500-2500"
	case -8:
		return "dutycycle outside set range"
	case -9:
		return "timer not 0-9"
	case -10:
		return "ms not 10-60000"
	case -11:
		return "timetype not 0-1"
	case -12:
		return "seconds < 0"
	case -13:
		return "micros not 0-999999"
	case -14:
		return "gpioSetTimerFunc failed"
	case -15:
		return "timeout not 0-60000"
	case -17:
		return "clock peripheral not 0-1"
	case -19:
		return "clock micros not 1, 2, 4, 5, 8, or 10"
	case -20:
		return "buf millis not 100-10000"
	case -21:
		return "dutycycle range not 25-40000"
	case -22:
		return "signum not 0-63"
	case -23:
		return "can't open pathname"
	case -24:
		return "no handle available"
	case -25:
		return "unknown handle"
	case -26:
		return "ifFlags > 4"
	case -27:
		return "Dmachannel not 0-15"
	case -28:
		return "socket port not 1024-32000"
	case -29:
		return "unrecognized fifo command"
	case -30:
		return "Dmasecondary channel not 0-15"
	case -31:
		return "function called before gpioInitialise"
	case -32:
		return "function called after gpioInitialise"
	case -33:
		return "waveform mode not 0-3"
	case -34:
		return "bad parameter in gpioCfgInternals call"
	case -35:
		return "baud rate not 50-250K(RX)/50-1M(TX)"
	case -36:
		return "waveform has too many pulses"
	case -37:
		return "waveform has too many chars"
	case -38:
		return "no bit bang serial read on GPIO"
	case -39:
		return "bad (null) serial structure parameter"
	case -40:
		return "bad (null) serial buf parameter"
	case -41:
		return "Gpiooperation not permitted"
	case -42:
		return "one or more Gpionot permitted"
	case -43:
		return "bad Wvscsubcommand"
	case -44:
		return "bad Wvsmsubcommand"
	case -45:
		return "bad Wvspsubcommand"
	case -46:
		return "trigger pulse length not 1-100"
	case -47:
		return "invalid script"
	case -48:
		return "unknown script id"
	case -49:
		return "add serial data offset > 30 minutes"
	case -50:
		return "Gpioalready in use"
	case -51:
		return "must read at least a byte at a time"
	case -52:
		return "script parameter id not 0-9"
	case -53:
		return "script has duplicate tag"
	case -54:
		return "script has too many tags"
	case -55:
		return "illegal script command"
	case -56:
		return "script variable id not 0-149"
	case -57:
		return "no more room for scripts"
	case -58:
		return "can't allocate temporary memory"
	case -59:
		return "socket read failed"
	case -60:
		return "socket write failed"
	case -61:
		return "too many script parameters (> 10)"
	case -62:
		return "script initialising"
	case -63:
		return "script has unresolved tag"
	case -64:
		return "bad Micsdelay (too large)"
	case -65:
		return "bad Milsdelay (too large)"
	case -66:
		return "non existent wave id"
	case -67:
		return "No more CBs for waveform"
	case -68:
		return "No more Oolfor waveform"
	case -69:
		return "attempt to create an empty waveform"
	case -70:
		return "no more waveforms"
	case -71:
		return "can't open I2C device"
	case -72:
		return "can't open serial device"
	case -73:
		return "can't open Spidevice"
	case -74:
		return "bad I2C bus"
	case -75:
		return "bad I2C address"
	case -76:
		return "bad Spichannel"
	case -77:
		return "bad i2c/spi/ser open flags"
	case -78:
		return "bad Spispeed"
	case -79:
		return "bad serial device name"
	case -80:
		return "bad serial baud rate"
	case -81:
		return "bad i2c/spi/ser parameter"
	case -82:
		return "i2c write failed"
	case -83:
		return "i2c read failed"
	case -84:
		return "bad Spicount"
	case -85:
		return "ser write failed"
	case -86:
		return "ser read failed"
	case -87:
		return "ser read no data available"
	case -88:
		return "unknown command"
	case -89:
		return "spi xfer/read/write failed"
	case -90:
		return "bad (NULL) pointer"
	case -91:
		return "no auxiliary Spion Pi A or B"
	case -92:
		return "Gpiois not in use for GpioPwm"
	case -93:
		return "Gpiois not in use for servo pulses"
	case -94:
		return "Gpiohas no hardware clock"
	case -95:
		return "Gpiohas no hardware GpioPwm"
	case -96:
		return "invalid hardware Pwmfrequency"
	case -97:
		return "hardware Pwmdutycycle not 0-1M"
	case -98:
		return "invalid hardware clock frequency"
	case -99:
		return "need password to use hardware clock 1"
	case -100:
		return "illegal, Pwmin use for main clock"
	case -101:
		return "serial data bits not 1-32"
	case -102:
		return "serial (half) stop bits not 2-8"
	case -103:
		return "socket/pipe message too big"
	case -104:
		return "bad memory allocation mode"
	case -105:
		return "too many I2C transaction segments"
	case -106:
		return "an I2C transaction segment failed"
	case -107:
		return "SMBus command not supported by driver"
	case -108:
		return "no bit bang I2C in progress on GPIO"
	case -109:
		return "bad I2C write length"
	case -110:
		return "bad I2C read length"
	case -111:
		return "bad I2C command"
	case -112:
		return "bad I2C baud rate, not 50-500k"
	case -113:
		return "bad chain loop count"
	case -114:
		return "empty chain loop"
	case -115:
		return "too many chain counters"
	case -116:
		return "bad chain command"
	case -117:
		return "bad chain delay micros"
	case -118:
		return "chain counters nested too deeply"
	case -119:
		return "chain is too long"
	case -120:
		return "deprecated function removed"
	case -121:
		return "bit bang serial invert not 0 or 1"
	case -122:
		return "bad Isredge value, not 0-2"
	case -123:
		return "bad Isrinitialisation"
	case -124:
		return "loop forever must be last command"
	case -125:
		return "bad filter parameter"
	case -126:
		return "bad pad number"
	case -127:
		return "bad pad drive strength"
	case -128:
		return "file open failed"
	case -129:
		return "bad file mode"
	case -130:
		return "bad file flag"
	case -131:
		return "bad file read"
	case -132:
		return "bad file write"
	case -133:
		return "file not open for read"
	case -134:
		return "file not open for write"
	case -135:
		return "bad file seek"
	case -136:
		return "no files match pattern"
	case -137:
		return "no permission to access file"
	case -138:
		return "file is a directory"
	case -139:
		return "bad shell return status"
	case -140:
		return "bad script name"
	case -141:
		return "bad Spibaud rate, not 50-500k"
	case -142:
		return "no bit bang Spiin progress on GPIO"
	case -143:
		return "bad event id"
	case -144:
		return "Used by Python"
	case -145:
		return "not available on BCM2711"
	case -146:
		return "only available on BCM2711"
	default:
		return "unknown code"
	}
}
