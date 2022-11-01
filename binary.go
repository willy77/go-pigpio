package GoPiGPIO

import (
	"bytes"
)

func convertToInt16(data []byte) int16 {
	switch len(data) {
	case 0:
		return 0
	case 1:
		return int16(data[0])
	default:
		return int16(data[0]) | (int16(data[1]) << 8)
	}
}

func convertToUint16(data []byte) uint16 {
	switch len(data) {
	case 0:
		return 0
	case 1:
		return uint16(data[0])
	default:
		return uint16(data[0]) | (uint16(data[1]) << 8)
	}
}

func convertToInt32(data []byte) int {
	switch len(data) {
	case 0:
		return 0
	case 1:
		return int(data[0])
	case 2:
		return int(int32(data[0]) | (int32(data[1]) << 8))
	case 3:
		return int(int32(data[0]) | (int32(data[1]) << 8) |
			(int32(data[2]) << 16))
	default:
		return int(int32(data[0]) | (int32(data[1]) << 8) |
			(int32(data[2]) << 16) | (int32(data[3]) << 24))
	}
}

func convertToUint32(data []byte) uint {
	switch len(data) {
	case 0:
		return 0
	case 1:
		return uint(data[0])
	case 2:
		return uint(uint32(data[0]) | (uint32(data[1]) << 8))
	case 3:
		return uint(uint32(data[0]) | (uint32(data[1]) << 8) |
			(uint32(data[2]) << 16))
	default:
		return uint(uint32(data[0]) | (uint32(data[1]) << 8) |
			(uint32(data[2]) << 16) | (uint32(data[3]) << 24))
	}
}

func convertToInt16Array(data []byte) []int16 {
	arr := make([]int16, 0)
	for offset := 0; offset < len(data); offset += 2 {
		arr = append(arr, convertToInt16(data[offset:]))
	}

	return arr
}

func convertToUint16Array(data []byte) []uint16 {
	arr := make([]uint16, 0)
	for offset := 0; offset < len(data); offset += 2 {
		arr = append(arr, convertToUint16(data[offset:]))
	}

	return arr
}

func convertToInt32Array(data []byte) []int {
	arr := make([]int, 0)
	for offset := 0; offset < len(data); offset += 4 {
		arr = append(arr, convertToInt32(data[offset:]))
	}

	return arr
}

func convertToUint32Array(data []byte) []uint {
	arr := make([]uint, 0)
	for offset := 0; offset < len(data); offset += 4 {
		arr = append(arr, convertToUint32(data[offset:]))
	}

	return arr
}

func convertToBytes(v ...int) []byte {
	buf := new(bytes.Buffer)
	for _, val := range v {
		buf.WriteByte(byte(val))
		buf.WriteByte(byte(val >> 8))
		buf.WriteByte(byte(val >> 16))
		buf.WriteByte(byte(val >> 24))
	}

	return buf.Bytes()
}
