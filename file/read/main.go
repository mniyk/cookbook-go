package main

import (
	"fmt"
	"os"
)

func main() {
	d, err := os.ReadFile("../write/output.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	fmt.Println(string(d))
}
