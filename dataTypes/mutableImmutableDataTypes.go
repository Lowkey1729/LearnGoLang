package main

import "fmt"

func main() {
	x := map[string]string{
		"name":  "Mojeed",
		"woman": "Fadheelah",
	}

	y := x
	y["married_to_her_yet"] = "Insha Allah"

	fmt.Println(x, y)
}
