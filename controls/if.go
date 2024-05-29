package main

import (
	"fmt"
	"math/rand"
)

func isOdd(number int) bool {

	if number%2 == 1 {
		return true
	}
	return false
}

func superIf() {

	// if 블록에만 쓸 변수 선언
	if randomNumber := rand.Intn(10); randomNumber < 5 {
		fmt.Println("통과")
		return
	}

	// randomNumber -> 밖에서는 못 쓴다.
	fmt.Println("탈!락!")
}

func superIf2() {

	// if 블록에만 쓸 변수 선언
	if randomNumber := rand.Intn(10); randomNumber < 5 {
		fmt.Println("통과")
		return
	} else {

		// else 블럭에선 쓸 수 있다.
		fmt.Println("탈락 해부렀으 ", randomNumber)
	}
}
