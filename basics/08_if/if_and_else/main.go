package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // 선언한 변수는 else에서도 사용가능
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
func main()  {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}