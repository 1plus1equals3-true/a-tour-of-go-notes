package main

import "fmt"

func main() {
	// 다른 함수가 실행될 때까지 해당 함수 지연
	// 인자는 즉시 평가되지만(z), 함수 호출은 지연
	z := 1
	defer fmt.Println("world", z)
	z += z
	fmt.Println("hello", z)
}
