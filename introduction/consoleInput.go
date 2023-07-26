package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter your name: ")

	scanner.Scan()

	input := scanner.Text()

	fmt.Printf("You entered: %s", input)

	//calculate my age

	fmt.Printf("Enter your birth year: ")

	scanner.Scan()

	nextInput, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("You will be %d years old  at the end of the year 2023", 2023-nextInput)

}
