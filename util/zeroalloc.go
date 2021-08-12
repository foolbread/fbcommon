/*
author: foolbread
file: zeroalloc.go
date: 2020/10/13 11:27 上午
*/
package util

import "unsafe"

func StringToByteSlice(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*unsafe.Sizeof(&b))) = len(s)
	return
}

func ByteSliceToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Uint64ToInt64(val uint64) int64 {
	return *(*int64)(unsafe.Pointer(&val))
}

func Uint64ToFloat64(val uint64) float64 {
	return *(*float64)(unsafe.Pointer(&val))
}

func Int64ToUint64(val int64) uint64 {
	return *(*uint64)(unsafe.Pointer(&val))
}

func Float64ToUint64(val float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&val))
}