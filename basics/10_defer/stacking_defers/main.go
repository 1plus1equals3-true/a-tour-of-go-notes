package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i<10; i++ {
		// 지연된 함수는 스택에 저장
		// 스택에 저장된 지연 함수는 후입선출로 반환
		defer fmt.Println(i)
	}

	fmt.Println("done")
}