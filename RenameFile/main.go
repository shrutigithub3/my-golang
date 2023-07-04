package main

import (
	"log"
	"os"
)

func main() {
	Original_Path := "myfile.txt"
	New_Path := "data.txt"
	e := os.Rename(Original_Path, New_Path)
	if e != nil {
		log.Fatal(e)
	}

}
