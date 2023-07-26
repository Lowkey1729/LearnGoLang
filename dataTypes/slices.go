package main

import "fmt"

func main() {

	x := [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:4]
	s = append(s, 100)

	length := fmt.Sprintf("\n The length of the slice is %d", len(s))

	capacity := fmt.Sprintf("\n The capacity of the slice is %d", cap(s))

	fmt.Println(s, length, capacity, s[:cap(s)])
}
