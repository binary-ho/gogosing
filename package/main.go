package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	N := 10
	randomNumber := rand.Intn(N)
	fmt.Println("%g 미만, 0 이상 랜덤 숫자 출력 : %g", N, randomNumber)

	// 대소문자 다 되는데??
	imports()
	Imports()

	fmt.Println(add(1, 2))
	fmt.Println(add2(1, 2, 3))

	// 복수 Return Test
	string1, string2 := swap("나는", "짱")
	fmt.Println(string1, string2)

	fmt.Println(getNamedResults(33))

	// 변수들
	getVariable()
	fmt.Println(Variable)
	fmt.Println(c, python, java)
	fmt.Println(initVar1, initVar2, initVar3)

	var floatValue = float32(math.Sqrt(7))
	fmt.Println("형변환은 쉬워 : ", convertFloatToInt(floatValue))
	getNumericConstants()
}
