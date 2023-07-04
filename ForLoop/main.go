package main

import (
	"fmt"
)

func main() {
	// var name = "shrutikumra"
	// var character = ""
	// fmt.Println("enter the character:")
	// fmt.Scanln(&name)
	// res1 := strings.Count(name, character)

	// fmt.Println("Character occured:", res1, "times")
	var character, count int
	count = 0
	name := "shrutikumra"
	fmt.Println("The string value is :", name)

	fmt.Print("Enter character = ")
	fmt.Scanln(&character)

	for character > 1 {
		count = count + 1
	}
	fmt.Println("The character occurs = ", count)
}
