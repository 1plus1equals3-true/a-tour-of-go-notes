package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // 조건없을때 switch true 동일
	case t.Hour() < 12 :
		fmt.Println("Good morning!")
	case t.Hour() < 17 :
		fmt.Println("Good afternoon.")
	default :
		fmt.Println("Good evening.")
	}
}