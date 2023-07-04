package main

import "fmt"

func main() {
	var a interface{}
	a = 10
	interfaceValue := a.(int)
	fmt.Println("Value is : ", interfaceValue)

}
