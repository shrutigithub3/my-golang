package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "myfile.txt"
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Successfully deleted file: ", fileName)
	}
}
