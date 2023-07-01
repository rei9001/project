package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[0] = Product{"임상현", 10000}

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, 0)

	v, ok := m[0]

	fmt.Println(v, ok)
}
