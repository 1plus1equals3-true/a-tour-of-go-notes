package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// 초기화 없이 선언된 변수는 모두 zero값을 갖는다 (0, false, "")