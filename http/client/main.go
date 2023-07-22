package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
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
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %d", err)
	}
	fmt.Println(string(b))
}
