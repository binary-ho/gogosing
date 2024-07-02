package main

import "fmt"

func testDefer() {

	// defer은 stack에 쌓이기 때문에 222가 먼저 출력된다.
	defer fmt.Println("내가 먼저야 111")
	defer fmt.Println("222")
	fmt.Println("333")
}
