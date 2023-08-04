package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// 매개변수 타입 생략 가능 (동일한 경우에만)
func addString(x, y string) string {
	return x + y
}

// 다중 반환 값 작성 가능
func multiReturn(x, y string) (string, string) {
	return x + y, y + x
}

// 이름으로 반환 값 작성 가능
func namedReturn(num int) (x, y int) {
	x = num % 2
	y = num / 2
	return // "naked" return
}

func main() {
	fmt.Println(add(100, 100))
	fmt.Println(addString("ab", "cd"))

	a, b := multiReturn("ab", "cd")
	fmt.Println(a, b)

	fmt.Println(namedReturn(105))
}

/*
- 함수는 0개 혹은 그보다 많은 인자를 받을 수 있다.

- 변수 이름 뒤에 type이 온다는 것을 명심해야한다.

- Go 진영에서 변수 이름 뒤에 type을 지정하도록 한 이유는 다음 글에서 찾아볼 수 있다. ( https://go.dev/blog/declaration-syntax )

- addString 함수와 같이 만일 함수 매개변수가 같은 type일 때 마지막 매개변수를 제외한 매개변수들의 type을 생략할 수 있다.

- multiReturn 함수와 같이 한 함수는 몇 개의 결과든 반환할 수 있다.

- namedReturn 함수와 같이 Go의 반환 값이 이름으로 정해질 수도 있다.
	- 이러한 이름은 반환 값의 의미를 설명하는데 사용되어야한다.
	- 인자가 없는 return 문은 이름이 주어진 반환 값을 반환한다. 이것을 "naked" return 이라고 한다.
	- "naked" return 문은 짧은 함수에서만 사용되어야한다. 긴 함수에서 사용하면 가독성을 떨어뜨릴 수 있다.
*/
