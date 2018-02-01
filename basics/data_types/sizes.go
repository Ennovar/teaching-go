package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

// The sizes of types int and uint (not their more specific counterparts) are
// determined at compile time by what is stored in $GOARCH. In other words,
// generally speaking, if the code is compiled on a 64-bit system, the sizes
// of type int and uint will be 64-bits. Likewise, if the code is compiled on
// a 32-bit system, they will both be 32-bits.

// If the specific types int64, uint64, or float64 are compiled by a 32-bit
// system, the compiler is smart enough to allocate two memory locations for
// them to utilize.

var intRef int
var int8Ref int8
var int16Ref int16
var int32Ref int32
var int64Ref int64

var uintRef uint
var uint8Ref uint8
var uint16Ref uint16
var uint32Ref uint32
var uint64Ref uint64

// Type byte is an alias for type uint8
var byteRef byte

var float32Ref float32
var float64Ref float64

var boolRef bool

var stringRef string

func main() {
	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(intRef), unsafe.Sizeof(intRef), unsafe.Sizeof(intRef)*8)
	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(int8Ref), unsafe.Sizeof(int8Ref), unsafe.Sizeof(int8Ref)*8)
	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(int16Ref), unsafe.Sizeof(int16Ref), unsafe.Sizeof(int16Ref)*8)
	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(int32Ref), unsafe.Sizeof(int32Ref), unsafe.Sizeof(int32Ref)*8)
	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(int64Ref), unsafe.Sizeof(int64Ref), unsafe.Sizeof(int64Ref)*8)

	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(uintRef), unsafe.Sizeof(uintRef), unsafe.Sizeof(uintRef)*8)
	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(uint8Ref), unsafe.Sizeof(uint8Ref), unsafe.Sizeof(uint8Ref)*8)
	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(uint16Ref), unsafe.Sizeof(uint16Ref), unsafe.Sizeof(uint16Ref)*8)
	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(uint32Ref), unsafe.Sizeof(uint32Ref), unsafe.Sizeof(uint32Ref)*8)
	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(uint64Ref), unsafe.Sizeof(uint64Ref), unsafe.Sizeof(uint64Ref)*8)

	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(byteRef), unsafe.Sizeof(byteRef), unsafe.Sizeof(byteRef)*8)

	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(float32Ref), unsafe.Sizeof(float32Ref), unsafe.Sizeof(float32Ref)*8)
	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(float64Ref), unsafe.Sizeof(float64Ref), unsafe.Sizeof(float64Ref)*8)

	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n\n", reflect.TypeOf(boolRef), unsafe.Sizeof(boolRef), unsafe.Sizeof(boolRef)*8)

	fmt.Printf("Type:  %v\nBytes: %v\nBits:  %v\n", reflect.TypeOf(stringRef), unsafe.Sizeof(stringRef), unsafe.Sizeof(stringRef)*8)
}
