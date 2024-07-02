package inittest

import "fmt"

type InitTest2 struct {
	foo, bar, ooo int
}

func init() {
	fmt.Println("inittest package func init()222222222222")
}
