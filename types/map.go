package main

import "fmt"

type Axis struct {
	y, x int
}

func testMap() {
	m := map[string]Axis{}
	m2 := make(map[string]Axis)

	m["axis1"] = Axis{7, 7}
	m2["axis2"] = Axis{777, 777}

	fmt.Println(m)
	fmt.Println("=======================")
	fmt.Println(m["axis1"])
	fmt.Println("=======================")
	fmt.Println(m2)
	fmt.Println("=======================")
	fmt.Println(m2["axis2"])
	fmt.Println("=======================")
	_, exists := m2["없는 Key인데 걍 기본값 뱉네"]
	fmt.Println(exists)
}
