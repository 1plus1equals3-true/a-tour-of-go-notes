package main

import "fmt"

func split(sum int) (x, y int) { // 리턴값의 타입뿐 아니라 이름까지 지정하면 return시 반영됨
	x = sum * 4 / 9
	y = sum - x
	return // 따로 뭘 리턴할지 정하지않는 형태 naked return
}

func main()  {
	fmt.Println(split(17))
}