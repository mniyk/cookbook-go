package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello World!")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}
