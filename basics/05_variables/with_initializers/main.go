package main

import "fmt"

var i, j int = 1, 2 // 타입을 포함하여 선언하고 초기화

func main() {
	var c, python, java = true, false, "no!" // 타입 없이 선언하고 초기화 (타입은 추론)
	fmt.Println(i, j, c, python, java)
}
