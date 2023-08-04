package main

// 1. "factored" import문 : 아래와 같이 import를 그룹 짓는 방법
import (
	"fmt"
	"math"
)

// 2. 물론 아래와 같이 다중 import문 작성도 가능
/**
import "fmt"
import "math"
*/

func main() {
	fmt.Printf("Now I have %g programs. \n", math.Sqrt(7))
	fmt.Printf("Now I have %g programs. \n", math.Sqrt(10))
}
