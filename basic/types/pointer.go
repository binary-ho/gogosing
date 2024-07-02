package main

import "fmt"

func testPointer() {

	i, j := 42, 2701
	p := &i
	fmt.Printf("p 주소 : %d, 값 : %d\n", p, *p)

	*p = 21
	fmt.Println("new i : ", i)

	p = &j
	*p = *p / 77
	fmt.Println("new j : ", j)
}
