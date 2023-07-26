package main

import "fmt"

func main() {
	//var mp map[string]int = map[string]int{
	//	"apple":  5,
	//	"pear":   6,
	//	"orange": 9,
	//}

	mp := map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	mp["pear"] = 900

	//delete(mp, "apple")

	val, ok := mp["apple"]

	if ok {
		println("Apple key is present", val)
	}

	fmt.Println(mp, mp["apple"])
}
