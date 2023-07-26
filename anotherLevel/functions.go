package main

import "fmt"

func main() {
	whoIsShe("Fadheelah \n")
	whoIsShe("Morenikeji \n")
	whoIsShe("Bolaji \n")

	herName, herSurname := herNAme("Fadheelah")

	firstName, middleName := herOtherNames("Morenikeji", "AbdulMojeed")

	fmt.Printf("%s %s \n", herName, herSurname)
	fmt.Printf("%s %s \n", firstName, middleName)
}

func whoIsShe(name string) {
	fmt.Printf("She is %s", name)
}

func herNAme(name string) (string, string) {
	return name, "Yusuf"
}

func herOtherNames(name1, name2 string) (firstName string, middleName string) {
	firstName = name1
	middleName = name2
	return
}
