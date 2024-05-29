package main

import "fmt"

func goLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 10
	}
	fmt.Println(sum)
}

// MEMO : go의 for은 초기화 구문이나 사후 구문을 생략해도 된다.
// 이 형태의 for은 while을 대체한다
func goLoop2() {
	sum := 0
	for sum <= 100 {
		sum += 10
	}
	fmt.Println(sum)
}
