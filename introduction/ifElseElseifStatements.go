package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your first name: ")

	scanner.Scan()

	firstName := scanner.Text()

	fmt.Print("Enter your second name: ")

	scanner.Scan()

	secondName := scanner.Text()

	if firstName == secondName {
		fmt.Println("You entered the same name!")
	}
}
