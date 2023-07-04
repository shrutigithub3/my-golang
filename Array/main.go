package main

import "fmt"

func main() {

	var a [4]string
	a[0] = "Shruti"
	a[1] = "Shiffali"
	a[2] = "Mansi"
	a[3] = "Anshu"
	fmt.Println("Names list is : ", a)

	for i := 0; i < 2; i++ {
		fmt.Println("i", i)
		fmt.Println("Elements are : ", a[i])
	}
}
