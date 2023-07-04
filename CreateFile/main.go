package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Content := "HELLO,SHRUTI HERE"
	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, Content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length is ", length)

	defer file.Close()
	fmt.Println("File created successfully")
}
func readFile() {
	content, err := os.ReadFile("myfile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
