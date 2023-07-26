package main

import "fmt"

func main() {
	a := []int{100, 120, 980, 540, 200, 56, 87, 123, 567, 567}

	//for i := 0; i < len(a); i++ {
	//	fmt.Println(a[i])
	//}

	//for _, element := range a {
	//	fmt.Println(element)
	//}

	// print an element that has duplicate once

	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}
}
