package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //if문 조건 시작할때 변수 선언 가능 v := math.Pow(x, n);
		return v
	}
	return lim
}

func main()  {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}