//go:build go1.16

package bits

import (
	"reflect"
	"unsafe"
)

func BoolToBytes(data []bool) []byte {
	// return unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data)
	header.Cap = len(data)
	bytes := *(*[]byte)(unsafe.Pointer(&header))
	return bytes

}

func Int8ToBytes(data []int8) []byte {
	// return unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data)
	header.Cap = len(data)
	bytes := *(*[]byte)(unsafe.Pointer(&header))
	return bytes
}

func Int16ToBytes(data []int16) []byte {
	// return unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), 2*len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = 2 * len(data)
	header.Cap = 2 * len(data)
	bytes := *(*[]byte)(unsafe.Pointer(&header))
	return bytes
}

func Int32ToBytes(data []int32) []byte {
	// return unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), 4*len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = 4 * len(data)
	header.Cap = 4 * len(data)
	bytes := *(*[]byte)(unsafe.Pointer(&header))
	return bytes
}

func Int64ToBytes(data []int64) []byte {
	// return unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), 8*len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = 8 * len(data)
	header.Cap = 8 * len(data)
	bytes := *(*[]byte)(unsafe.Pointer(&header))
	return bytes
}

func Float32ToBytes(data []float32) []byte {
	// return unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), 4*len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = 4 * len(data)
	header.Cap = 4 * len(data)
	bytes := *(*[]byte)(unsafe.Pointer(&header))
	return bytes
}

func Float64ToBytes(data []float64) []byte {
	// return unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), 8*len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = 8 * len(data)
	header.Cap = 8 * len(data)
	bytes := *(*[]byte)(unsafe.Pointer(&header))
	return bytes
}

func Int16ToUint16(data []int16) []uint16 {
	// return unsafe.Slice(*(**uint16)(unsafe.Pointer(&data)), len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data)
	header.Cap = len(data)
	uint16s := *(*[]uint16)(unsafe.Pointer(&header))
	return uint16s
}

func Int32ToUint32(data []int32) []uint32 {
	// return unsafe.Slice(*(**uint32)(unsafe.Pointer(&data)), len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data)
	header.Cap = len(data)
	uint32s := *(*[]uint32)(unsafe.Pointer(&header))
	return uint32s
}

func Int64ToUint64(data []int64) []uint64 {
	// return unsafe.Slice(*(**uint64)(unsafe.Pointer(&data)), len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data)
	header.Cap = len(data)
	uint64s := *(*[]uint64)(unsafe.Pointer(&header))
	return uint64s
}

func Float32ToUint32(data []float32) []uint32 {
	// return unsafe.Slice(*(**uint32)(unsafe.Pointer(&data)), len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data)
	header.Cap = len(data)
	uint32s := *(*[]uint32)(unsafe.Pointer(&header))
	return uint32s
}

func Float64ToUint64(data []float64) []uint64 {
	// return unsafe.Slice(*(**uint64)(unsafe.Pointer(&data)), len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data)
	header.Cap = len(data)
	uint64s := *(*[]uint64)(unsafe.Pointer(&header))
	return uint64s
}

func Uint32ToBytes(data []uint32) []byte {
	// return unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), 4*len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = 4 * len(data)
	header.Cap = 4 * len(data)
	bytes := *(*[]byte)(unsafe.Pointer(&header))
	return bytes
}

func Uint64ToBytes(data []uint64) []byte {
	// return unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), 8*len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = 8 * len(data)
	header.Cap = 8 * len(data)
	bytes := *(*[]byte)(unsafe.Pointer(&header))
	return bytes
}

func Uint128ToBytes(data [][16]byte) []byte {
	// return unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), 16*len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = 16 * len(data)
	header.Cap = 16 * len(data)
	bytes := *(*[]byte)(unsafe.Pointer(&header))
	return bytes
}

func Uint32ToInt32(data []uint32) []int32 {
	// return unsafe.Slice(*(**int32)(unsafe.Pointer(&data)), len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data)
	header.Cap = len(data)
	int32s := *(*[]int32)(unsafe.Pointer(&header))
	return int32s
}

func Uint64ToInt64(data []uint64) []int64 {
	// return unsafe.Slice(*(**int64)(unsafe.Pointer(&data)), len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data)
	header.Cap = len(data)
	int64s := *(*[]int64)(unsafe.Pointer(&header))
	return int64s
}

func BytesToBool(data []byte) []bool {
	// return unsafe.Slice(*(**bool)(unsafe.Pointer(&data)), len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data)
	header.Cap = len(data)
	bools := *(*[]bool)(unsafe.Pointer(&header))
	return bools
}

func BytesToInt8(data []byte) []int8 {
	// return unsafe.Slice(*(**int8)(unsafe.Pointer(&data)), len(data))
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data)
	header.Cap = len(data)
	int8s := *(*[]int8)(unsafe.Pointer(&header))
	return int8s
}

func BytesToInt16(data []byte) []int16 {
	// return unsafe.Slice(*(**int16)(unsafe.Pointer(&data)), len(data)/2)
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data) / 2
	header.Cap = len(data) / 2
	int16s := *(*[]int16)(unsafe.Pointer(&header))
	return int16s
}

func BytesToInt32(data []byte) []int32 {
	// return unsafe.Slice(*(**int32)(unsafe.Pointer(&data)), len(data)/4)
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data) / 4
	header.Cap = len(data) / 4
	int32s := *(*[]int32)(unsafe.Pointer(&header))
	return int32s
}

func BytesToInt64(data []byte) []int64 {
	// return unsafe.Slice(*(**int64)(unsafe.Pointer(&data)), len(data)/8)
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data) / 8
	header.Cap = len(data) / 8
	int64s := *(*[]int64)(unsafe.Pointer(&header))
	return int64s
}

func BytesToUint32(data []byte) []uint32 {
	// return unsafe.Slice(*(**uint32)(unsafe.Pointer(&data)), len(data)/4)
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data) / 4
	header.Cap = len(data) / 4
	uint32s := *(*[]uint32)(unsafe.Pointer(&header))
	return uint32s
}

func BytesToUint64(data []byte) []uint64 {
	// return unsafe.Slice(*(**uint64)(unsafe.Pointer(&data)), len(data)/8)
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data) / 8
	header.Cap = len(data) / 8
	uint64s := *(*[]uint64)(unsafe.Pointer(&header))
	return uint64s
}

func BytesToUint128(data []byte) [][16]byte {
	// return unsafe.Slice(*(**[16]byte)(unsafe.Pointer(&data)), len(data)/16)
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data) / 16
	header.Cap = len(data) / 16
	x := *(*[][16]byte)(unsafe.Pointer(&header))
	return x
}

func BytesToFloat32(data []byte) []float32 {
	// return unsafe.Slice(*(**float32)(unsafe.Pointer(&data)), len(data)/4)
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data) / 4
	header.Cap = len(data) / 4
	float32s := *(*[]float32)(unsafe.Pointer(&header))
	return float32s
}

func BytesToFloat64(data []byte) []float64 {
	// return unsafe.Slice(*(**float64)(unsafe.Pointer(&data)), len(data)/8)
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Len = len(data) / 8
	header.Cap = len(data) / 8
	float64s := *(*[]float64)(unsafe.Pointer(&header))
	return float64s
}

func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
