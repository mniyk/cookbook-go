package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.WriteFile("output.txt", []byte("Hello World!"), 0644)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}
