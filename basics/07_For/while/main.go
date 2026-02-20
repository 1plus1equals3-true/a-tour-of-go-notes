package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { // while이 없지만 init, post와 ;까지 다 생략하고 조건만 사용하여 for로 구현
		sum += sum
	}
	fmt.Println(sum)
}