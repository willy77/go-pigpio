package pigpio

// CallbackFunc is used for gpio callbacks
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

func (c *Callback) Handle() int { return c.handle }
func (c *Callback) Edge() Edge  { return c.edge }
func (c *Callback) Bit() int    { return c.bit }
