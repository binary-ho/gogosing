package main

import "fmt"

func testFor() {
	testLoopIteration()
}

// 이런 형태의 순회도 있다!!
func testLoopIteration() {
	array := []int{10, 20, 30, 40, 50}
	for index, value := range array {
		fmt.Println(index, value)
	}

	// 안 쓰려면 아래 Bar
	for _, value := range array {
		fmt.Println(value)
	}
}

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
