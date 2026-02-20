package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe	bool		= false
	MaxInt	uint64		= 1<<64 -1
	z		complex128	= cmplx.Sqrt(-5 + 12i)
) // 변수 선언도 임포트 처럼 블럭으로 분리 가능

func main()  {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// Go의 모든 베이직 타입

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128