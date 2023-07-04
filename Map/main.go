package main

import "fmt"

func main() {
	// a := map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	// fmt.Println("Car is: ", a)

	a := make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	fmt.Println("Car is: ", a)
	//Accessing map element
	fmt.Println(a["brand"])
	//Removing map element
	delete(a, "model")
	fmt.Println(a)
}
