package main

import "fmt"

func main() {
	var i, j int = 1, 2 	// 타입 포함 선언 및 초기화
	k := 3 					// var k = 3 단축 선언
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
