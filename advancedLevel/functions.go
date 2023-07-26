package main

import "fmt"

func main() {
	//test := func(x int) int {
	//	return x * -1
	//}(8)

	test := func(x int) int {
		return x * -1
	}

	test2(test)

	returnFunction("Yup yup")()

	fmt.Println(test(8))

}

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(3))
}

func returnFunction(x string) func() {
	return func() {
		fmt.Println(x)
	}
}
