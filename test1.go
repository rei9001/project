package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create file")
		return
	}

	defer fmt.Println("finally called")
	defer f.Close()
	defer fmt.Println("file close")

	fmt.Println("write hello world to file")
	fmt.Fprintln(f, "hello world")
}
