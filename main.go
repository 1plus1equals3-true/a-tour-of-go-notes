package main // 라이브러리가 아닌 실행 가능 파일이라고 선언

import "fmt" // format의 약자,  go는 깡통이라 이런것도 임포트해야함

func main() { // 진입점은 main 고정, public static void main(String[] args)
    fmt.Println("Hello, World!") // println 대소문자 구분으로 public, private구분
}