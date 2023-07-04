package main

import "fmt"

func main() {
	var a, b, c, d, e int

	fmt.Print("\nEnter five Numbers to find Largest = ")
	fmt.Scanln(&a, &b, &c, &d, &e)

	if a > b && a > c && a > d && a > e {
		fmt.Println(a, " is Greater")
	} else if b > a && b > c && b > d && b > e {
		fmt.Println(b, " is Greater")
	} else if c > a && c > b && c > d && c > e {
		fmt.Println(c, " is Greater")
	} else if d > a && d > b && d > c && d > e {
		fmt.Println(d, " is Greater")
	} else {
		fmt.Println("e is Greater ")
	}
}
