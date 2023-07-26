package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//enter the guessed number
	fmt.Println("Enter the guessed number: ")

	scanner.Scan()

	guessedNumber, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	for guessedNumber != 19 {
		fmt.Printf("You entered the incorrect numer, %d. \n Please, try again \n", guessedNumber)

		fmt.Println("Enter the guessed number: ")

		scanner.Scan()

		guessedNumber, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	}

	fmt.Printf("Congratulations!!!! \n %d is the right number. You guessed correctly!", guessedNumber)

}
