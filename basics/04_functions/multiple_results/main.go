package main

import "fmt"

func swap(x, y string) (string, string) {
	// x, y string = x string, y string
	// 리턴값 두개일 때 둘다 타입 선언 func swap() (string, string) {
	return y, x // 리턴값 여러개 가능
}

func main() {
	// 변수 단축 선언 a := 값 (var a = 값) 타입은 값으로 추론
	// 타입 명시도 가능 var a int
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}