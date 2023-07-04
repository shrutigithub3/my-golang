package main

import (
	"errors"
	"fmt"
)

func main() {

	message := "Hello"
	myError := errors.New("WRONG MESSAGE")

	if message != "Shruti" {
		fmt.Println(myError)
	}

}
