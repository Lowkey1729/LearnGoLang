package main

import "fmt"

func main() {
	changedValue := "Yes, Boss"

	fmt.Println(changedValue)

	changeValue(&changedValue)

	fmt.Println(changedValue)

	changeValue1(changedValue)

	fmt.Println(changedValue)
}

func changeValue(str *string) {
	*str = "Fadheelah!!"
}

func changeValue1(str string) {
	str = "Fadheelah Changed"
}
