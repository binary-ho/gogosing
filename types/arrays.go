package main

import (
	"fmt"
)

func testArrays() {
	var slice []int
	slice = append(slice, 7)
	slice = append(slice, 7)
	slice = append(slice, 7)
	slice = append(slice, 7)
	slice2 := append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice2)
}

func testMatrix() {
	mapp := [][]string{
		[]string{"0", "0", "0"},
		[]string{"0", "0", "0"},
		[]string{"0", "0", "0"},
	}

	mapp[0][0] = "1"
	mapp[2][2] = "2"
	mapp[1][2] = "1"
	mapp[1][0] = "2"
	mapp[0][2] = "3"

	fmt.Printf("%s\n", mapp[0])
}

func emptySliceLenTest() {
	var array []int
	fmt.Println(len(array))
}

func testSlice() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	var subNumbers = numbers[1:4]
	fmt.Println(subNumbers)
}

// 슬라이스는 래퍼런스다!!
func testIsSliceReference() {
	numbers := [6]int{1, 2, 3, 4, 5, 6}
	var subNumbers = numbers[1:4]
	subNumbers[0] = 999999999
	fmt.Println(numbers)
}
