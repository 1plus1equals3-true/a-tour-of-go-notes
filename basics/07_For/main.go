package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { // for문 뒤에 조건에 ()가 없다
		sum += i
		// go의 반복문은 for만 존재
	}// {}는 항상 필요
	fmt.Println(sum)
}

// for ; sum < 1000; { sum += sum } init, post는 없어도 된다 (조건만 있으면 돌아간다)