package main

import "fmt"

func main() {
	var a = 0
	if a%2 == 0 {
		fmt.Println("Number is even")
	} else if a == a {
		fmt.Println("Whole number")
	} else {
		fmt.Println("number")
	}

}
