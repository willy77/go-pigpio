package GoPiGPIO

type CallbackFunc func(*GpioPin, uint)

type Edge int8

const (
	EdgeRising  Edge = 0
	EdgeFalling Edge = 1
	EdgeEither  Edge = 2
)

type Callback struct {
	handle int
	edge   Edge
	bit    int
	fn     CallbackFunc
}
