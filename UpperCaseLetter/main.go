package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hi! how are you?"
	title := strings.Title(str)
	fmt.Println(title)
}
