package main

import "fmt"

const (
	// 1bit를 왼쪽으로 100자리 시프트 = 1뒤에 0이 100개인 이진수
	Big = 1 << 100

	// Big에서 오른쪽으로 99자리 시프트 = 10 = 2
	Small = Big >> 99
)

// 타입이 정해지지 않은 상수는 컨텍스트에서 필요한 타입으로 취급
func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big)) 불가능 int는 최대 64bit
}