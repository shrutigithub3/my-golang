package main

import "fmt"

func main() {
	mySlice := [4]string{"Shruti", "Shiffali", "Anshu", "Mansi"}
	for i, item := range mySlice {
		fmt.Println("Value of mySlice is: ", i, item)
	}
}
