package main

import "fmt"

// 상수는 := 를 사용할 수 없다.
const Pi = 3.14

func main() {
	const World = "세계"
	fmt.Println("Hello", World) 	// + 붙이기
	fmt.Println("Happy", Pi, "Day") // , 띄우기

	const Truth = true
	fmt.Println("Go rules?", Truth)
}