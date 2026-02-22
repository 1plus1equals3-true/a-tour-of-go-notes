package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os { // go는 해당하는 case만 실행하고 끝남 break필요없음
	case "darwin":					// case가 상수가 아니어도 되고 관련값이 정수가 아니어도 됨
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}