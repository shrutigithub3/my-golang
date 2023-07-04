package main

import "fmt"

func main() {
	var a, b, x, celsius, fahrenheit int
	fmt.Println("Press 1 For Convert Fahrenheit To Celsius\n")
	fmt.Println("Press 2 For Convert Celsius To Fahrenheit\n")
	fmt.Println("\nEnter Your Choice: ")

	fmt.Scanln("%d", &x)

	switch x {
	case 1:
		fmt.Println("\nEnter The Temperature in Fahrenheit: ")
		fmt.Scanln("%f", &a)
		celsius = 5 * (a - 32) / 9
		fmt.Println("\n\nCelsius Temperature is: %f ", celsius)
		break
	case 2:
		fmt.Println("\nEnter The Temperature in Celsius: ")
		fmt.Scanln("%f", &b)
		fahrenheit = ((9 * b) / 5) + 32
		fmt.Println("\n\nFahrenheit Temperature is: %f ", fahrenheit)
		break
	default:
		fmt.Println("\n\nWrong Choice.....Try Again!!!\n")
	}
	// x1 := 0
	// return
}
