package main

import "fmt"

func main() {
	myAge := 20

	fmt.Println("Value is :", myAge)
	ptr := &myAge
	fmt.Println("Value of ptr is :", ptr)
}
