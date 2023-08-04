package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello", math.Pi)
	// fmt.Println("hello", math.pi) // 오류
}

/*
- Go에서는 대문자로 시작하는 이름이 export 된다. 예를 들어 Pi 가 math 패키지에서 export 되어 사용되었듯이 pi는 대문자로 시작하지 않으므로 export되지 않는다.

- package를 import할 때, 그 패키지의 export된 이름만 참조할 수 있다.

- export되지 않은 이름들은 패키지 밖에서 접근할 수 없다.
*/
