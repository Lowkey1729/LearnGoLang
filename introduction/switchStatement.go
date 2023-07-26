package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	println("Enter your gender")

	scanner.Scan()

	input := scanner.Text()

	switch input {
	case "Boy", "boy":
		fmt.Println("You entered a boy")
	case "Girl", "girl":
		fmt.Println("You entered a girl")
	default:
		fmt.Println("No gender selected")

	}
}
