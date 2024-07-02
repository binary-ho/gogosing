package main

import "fmt"

var c, python, java bool
var Variable int

var initNumber1, initNumber2, initNumber3 int = 111, 222, 333
var initVar1, initVar2, initVar3 = 11, "initVar2", false

// MEMO : type은 여기 https://go-tour-ko.appspot.com/basics/11

const CONST_VALUE = "CONST_VALUE"

func getVariable() {
	var i int
	shortVarInit := "':='는 함수 밖에서는 사용할 수 없음"
	fmt.Println(i, c, python, java)
	fmt.Println(shortVarInit)
	fmt.Println(CONST_VALUE)
}

func convertFloatToInt(floatValue float32) int {
	return int(floatValue)
}
