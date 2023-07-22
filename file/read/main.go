package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	d, err := ioutil.ReadFile("../write/output.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Println(string(d))
}
