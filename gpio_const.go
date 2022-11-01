package GoPiGPIO

type GpioLevel int

// GpioLevels
const (
	Low   GpioLevel = 0
	High  GpioLevel = 1
	Off   GpioLevel = 0
	On    GpioLevel = 1
	Clear GpioLevel = 0
	Set   GpioLevel = 1
)

type GpioMode int

// GpioModes
const (
	ModeInput  GpioMode = 0
	ModeOutput GpioMode = 1
	ModeAlt0   GpioMode = 4
	ModeAlt1   GpioMode = 5
	ModeAlt2   GpioMode = 6
	ModeAlt3   GpioMode = 7
	ModeAlt4   GpioMode = 3
	ModeAlt5   GpioMode = 2
)

type PullUpDownMode int

// Pull-Up-Down
const (
	PudOff  PullUpDownMode = 0
	PudDown PullUpDownMode = 1
	PudUp   PullUpDownMode = 2
)
