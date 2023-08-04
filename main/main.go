package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(
		"Study GoLang !! Like",
		rand.Intn(10),
		"times!",
	)
}

/*
- 모든 Go 프로그램은 패키지로 구성되어있음.

- 프로그램은 *main 패키지*에서 실행을 시작

- 관습적으로, 패키지의 이름은 import 경로의 마지막 요소와 같음. 예를 들어 "math/rand" 패키지는 package rand 문으로 시작하는 파일들로 구성됨.
*/
