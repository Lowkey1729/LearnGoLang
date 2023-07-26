package main

import "fmt"

func main() {
	arr := [3]int{1, 10, 19}
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
