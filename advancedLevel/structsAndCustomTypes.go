package main

import "fmt"

func main() {
	p1 := &Point{1, 2}
	p2 := Point{4, 5}

	fmt.Println(p1, p2, p1.x, p1.y)
	alterPoint(p1)
	fmt.Println(p1)
}

type Point struct {
	x int32
	y int32
}

func alterPoint(point *Point) {
	point.x = 890
}
