package main

import "fmt"

func main() {
	var num, remainder int

	fmt.Print("Enter the Number to check Palindrome = ")
	fmt.Scanln(&num)

	reverse := 0

	for temp := num; temp > 0; temp = temp / 10 {
		remainder = temp % 10
		reverse = reverse*10 + remainder
	}

	fmt.Println("The Reverse of the Given Number = ", reverse)
	if num == reverse {
		fmt.Println(num, " is a Palindrome Number")
	} else {
		fmt.Println(num, " is Not a Palindrome Number")
	}
}
