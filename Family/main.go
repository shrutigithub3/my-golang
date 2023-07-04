package main

import (
	parent "Family/Father"
	child "Family/Father/Son"

	"fmt"
)

func main() {
	f := new(parent.Father)
	fmt.Println(f.Data("Mr. Jeremy Maclin"))

	c := new(child.Son)
	fmt.Println(c.Data("Riley Maclin"))
}
