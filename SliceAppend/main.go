package main

import "fmt"

func main() {
	var fruitList = []string{"Apple", "Mango", "Grapes"}
	fmt.Println("Fruits are: ", fruitList)
	fruitList = append(fruitList, "Banana")
	fmt.Println("Fruits are : ", fruitList)
	fruitList = append(fruitList[2:3])
	fmt.Println("Fruits are : ", fruitList)
}
