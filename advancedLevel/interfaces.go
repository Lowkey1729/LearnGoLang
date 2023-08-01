package main

type shape interface {
	area() float64
}
type circle struct {
	radius float64
}
type rectangle struct {
	width  float64
	height float64
}
