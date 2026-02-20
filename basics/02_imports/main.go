package main

import (
	"fmt" // 포맷
	"math"
)

func main()  {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	// 포맷 문자열 "Now you have %g problems. \n"
	// 서식 지정자 %g = 실수형 숫자의 간단한 형태 출력
	// 대체하는 값 math.Sqrt(7)

	// Printf함수 : 첫번째 인자는 포맷문자열, 그 뒤 인자들이 서식 지정자에 차례로 대체 (%g, %d, %s 등)
}