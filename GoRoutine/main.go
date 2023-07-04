package main

import (
	"fmt"
	"time"
)

func display(message string) {

	for {
		fmt.Println(message)

		time.Sleep(time.Second * 1)
	}
}

func main() {

	go display("Process 1")

	display("Process 2")

}
