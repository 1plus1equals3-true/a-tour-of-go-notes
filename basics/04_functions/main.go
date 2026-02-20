package main

import "fmt"

func add(x int, y int) int{
	// x int 변수명 먼저, 타입 뒤에
	// add(x, y int) 가능
	// add() int{} 리턴값의 타입
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}