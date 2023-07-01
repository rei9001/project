package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	p := Person{Name: "John", Age: 30, Email: "john@example.com"}

	b, err := json.Marshal(p)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
