package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("/newone/again", 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Directory hierarchy created successfully")
}
