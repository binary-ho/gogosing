package main

func getNamedResults(num int) (x, y int) {
	x = num * 1000
	y = num / 9999
	x = num * 9999
	return
}
