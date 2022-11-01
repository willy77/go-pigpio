package GoPiGPIO

type I2CCommand byte

// I2C commands
const (
	I2CCmdEnd         I2CCommand = 0
	I2CCmdEsc         I2CCommand = 1
	I2CCmdStart       I2CCommand = 2
	I2CCmdCombinedOn  I2CCommand = 2
	I2CCmdStop        I2CCommand = 3
	I2CCmdCombinedOff I2CCommand = 3
	I2CCmdAddress     I2CCommand = 4
	I2CCmdFlags       I2CCommand = 5
	I2CCmdRead        I2CCommand = 6
	I2CCmdWrite       I2CCommand = 7
)
