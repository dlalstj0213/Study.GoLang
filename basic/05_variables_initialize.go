package main

import "fmt"

var x, y int = 1, 2

// isTeacher := true // 오류: 패키지 내에서는 짧은 변수 선언 안됨

func main() {
	var name, gender, isStudent = "Minseo", "Male", false

	fmt.Println(x, y, name, gender, isStudent)

	isTeacher := true // 짧은 변수 선언
	fmt.Println(isTeacher)
}

/*
- 변수 선언은 한 변수 당 하나의 초기값을 포함할 수 있다.

- 만역 초기값이 존재한다면, type은 생략될 수 있다. 이러한 경우는 변수의 type이 초기값의 type으로 가지게 된다.

- *함수 내*에서는 := 라는 짧은 변수 선언은 암시적 type으로 var 선언처럼 사용될 수 있다.

- 함수 밖에서는 모든 선언이 키워드(var, func, ...)로 시작하므로 := 구문은 사용할 수 없다.
*/
