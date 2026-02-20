package main

import "fmt"

func main() {
	v := 42 // 명시적으로 타입을 선언하지 않으면 타입 추론
	fmt.Printf("v is of type %T\n", v)
}
