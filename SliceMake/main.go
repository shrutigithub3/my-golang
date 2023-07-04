package main

import "fmt"

func main() {
	a := make([]int, 5)
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4
	fmt.Println("a", a)

	b := make([]int, 0, 5)

	fmt.Println("b", b)

	c := b[:2]

	fmt.Println("c", c)

	d := c[2:5]

	fmt.Println("d", d)
}
