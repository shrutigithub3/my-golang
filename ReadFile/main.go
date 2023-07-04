package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("myfile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
