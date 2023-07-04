package main

import "fmt"

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {
	a := Address{
		Name:    "Shruti",
		city:    "Hoshiarpur",
		Pincode: 001234,
	}
	fmt.Println("Address is :", a)
	fmt.Println("Name is :", a.Name)
}
