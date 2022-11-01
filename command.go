package GoPiGPIO

type Command int

// Commands
const (
	cmdModeS             Command = 0
	cmdModeG             Command = 1
	cmdPUD               Command = 2
	cmdRead              Command = 3
	cmdWrite             Command = 4
	cmdPwmDutyCycleS     Command = 5
	cmdPwmRangeS         Command = 6
	cmdPwmFreqS          Command = 7
	cmdServoPWS          Command = 8
	cmdBank1R            Command = 10
	cmdBank2R            Command = 11
	cmdBank1C            Command = 12
	cmdBank2C            Command = 13
	cmdBank1S            Command = 14
	cmdBank2S            Command = 15
	cmdGetTick           Command = 16
	cmdHWVersion         Command = 17
	cmdNotifyBegin       Command = 19
	cmdNotifyClose       Command = 21
	cmdPwmRangeG         Command = 22
	cmdPwmFreqG          Command = 23
	cmdPwmRealRangeG     Command = 24
	cmdGetVersion        Command = 26
	cmdWaveClear         Command = 27
	cmdWaveAddGeneric    Command = 28
	cmdWaveAddSerial     Command = 29
	cmdWaveBusy          Command = 32
	cmdWaveHalt          Command = 33
	cmdWaveMicros        Command = 34
	cmdWavePulse         Command = 35
	cmdWaveCBS           Command = 36
	cmdTrigger           Command = 37
	cmdScript            Command = 38
	cmdScriptDelete      Command = 39
	cmdScriptRun         Command = 40
	cmdScriptStop        Command = 41
	cmdScriptStatus      Command = 45
	cmdWaveCreate        Command = 49
	cmdWaveDelete        Command = 50
	cmdWaveOnce          Command = 51
	cmdWaveRepeat        Command = 52
	cmdWaveNew           Command = 53
	cmdI2COpen           Command = 54
	cmdI2CClose          Command = 55
	cmdI2CReadDevice     Command = 56
	cmdI2CWriteDevice    Command = 57
	cmdI2CWriteQuick     Command = 58
	cmdI2CReadByte       Command = 59
	cmdI2CWriteByte      Command = 60
	cmdI2CReadByteData   Command = 61
	cmdI2CWriteByteData  Command = 62
	cmdI2CReadWordData   Command = 63
	cmdI2CWriteWordData  Command = 64
	cmdI2CReadBlockData  Command = 65
	cmdI2CWriteBlockData Command = 66
	cmdI2CProcessCall    Command = 69
	cmdSPIOpen           Command = 71
	cmdSPIClose          Command = 72
	cmdSPIRead           Command = 73
	cmdSPIWrite          Command = 74
	cmdSPIXfer           Command = 75
	cmdSerialOpen        Command = 76
	cmdSerialClose       Command = 77
	cmdSerialReadByte    Command = 78
	cmdSerialWriteByte   Command = 79
	cmdSerialRead        Command = 80
	cmdSerialWrite       Command = 81
	cmdSerialAvailable   Command = 82
	cmdPwmDutyCycleG     Command = 83
	cmdServoPWG          Command = 84
	cmdHWClock           Command = 85
	cmdHwPWM             Command = 86
	cmdI2CBatch          Command = 92
	cmdWaveChain         Command = 93
	cmdFilterGlitch      Command = 97
	cmdFilterNoise       Command = 98
	cmdNoIB              Command = 99
	cmdWaveTxUsingMode   Command = 100
	cmdWaveTxAt          Command = 101
	cmdEventM            Command = 115
	cmdScriptUpdate      Command = 117
	cmdWaveCAP           Command = 118
)
