package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func printVertex() {
	vertex := Vertex{1, 2}
	fmt.Println("before : ", vertex.Y, ", ", vertex.X)
	vertex.X = 10
	vertex.Y = 20
	fmt.Println("after : ", vertex.Y, ", ", vertex.X)
}

func printVertex2() {
	vertex := Vertex{1, 2}
	fmt.Println("before : ", vertex.Y, ", ", vertex.X)

	notPointerVertex := vertex
	notPointerVertex.Y = 20
	fmt.Println("after 1 : ", vertex.Y, ", ", vertex.X)
	fmt.Println("after 1 : ", notPointerVertex.Y, ", ", notPointerVertex.X)

	pointerVertex := &vertex
	pointerVertex.Y = 200
	fmt.Println("after 2 : ", vertex.Y, ", ", vertex.X)
}

func printVertex3() {
	vertex := Vertex{Y: 1, X: 2}
	fmt.Println("just print : ", vertex.Y, ", ", vertex.X)
}
