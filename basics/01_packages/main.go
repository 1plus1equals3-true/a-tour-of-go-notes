package main // 실행 파일

import (
	"fmt"       // 포맷
	"math/rand" // 랜덤
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10)) //임포트 경로의 마지막 사용 math/rand => rand
}
