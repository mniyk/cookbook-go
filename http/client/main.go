package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	url := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequestWithContext(
		context.Background(), http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: %d", res.StatusCode)
		os.Exit(1)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %d", err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}
