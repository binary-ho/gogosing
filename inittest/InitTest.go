package inittest

import "fmt"

type InitTest struct {
	foo, bar, ooo int
}

func (initTestStruct *InitTest) init() {
	fmt.Println("inittest package func (initTestStruct *InitTestStruct) init()")
}

func init() {
	fmt.Println("inittest package func init()")
}
