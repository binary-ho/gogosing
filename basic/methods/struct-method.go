package main

import "fmt"

type Axis struct {
	X, Y int
}

func (axis Axis) PrintAxis() {
	fmt.Printf("Before X : %d, Y : %d\n", axis.X, axis.Y)
	axis.X = 100
	axis.Y = 200
	fmt.Printf("After X : %d, Y : %d\n", axis.X, axis.Y)
}

// ChangeAxis
/* 포인터 리시버를 사용하는 데에는 두 가지 이유가 있습니다.
첫번째는, 메서드가 리시버가 가리키는 값을 수정할 수 있기 때문입니다.
두번째는 각각의 메서드 call에서의 value 복사 문제를 피하기 위해서입니다. */
func (axis *Axis) ChangeAxis(x int, y int) {
	axis.X = x
	axis.Y = y
}

func ChangeAxis2(axis *Axis, x int, y int) {
	axis.X = x
	axis.Y = y
}

func ChangeAxis3(axis Axis, x int, y int) {
	axis.X = x
	axis.Y = y
}

func printAxis() {
	axis := Axis{10, 20}
	axis.PrintAxis()
	fmt.Println("실제론 바뀌었나? -> ", axis.X, ", ", axis.Y)
	axis.ChangeAxis(99, 99)
	fmt.Println("이번엔? ", axis.X, ", ", axis.Y)

	ChangeAxis2(&axis, 1, 2)
	fmt.Println("22 이번엔? ", axis.X, ", ", axis.Y)

	ChangeAxis3(axis, 999999, 999999)
	fmt.Println("33 마지막? ", axis.X, ", ", axis.Y)
}
