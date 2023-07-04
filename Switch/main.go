package main

import "fmt"

func main() {
	var marks int
	fmt.Println("Enter marks of student:")
	fmt.Scanln(&marks)
	switch {
	case marks >= 90:
		fmt.Println("Your grade is A")
	case marks >= 80:
		fmt.Println("Your grade is B")
	case marks >= 70:
		fmt.Println("Your grade is C")
	case marks >= 60:
		fmt.Println("Your grade is D")
	case marks >= 50:
		fmt.Println("Your grade is E")
	case marks >= 40:
		fmt.Println("Your grade is F")
	default:
		fmt.Println("Not valid marks")
	}
}
