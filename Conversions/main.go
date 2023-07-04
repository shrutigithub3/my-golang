package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("Welcome to our shop")
	println("Plz rate us between 1 and 10")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	println("Thanks for giving rating,", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added number to the rating", numRating+1)
	}
}
