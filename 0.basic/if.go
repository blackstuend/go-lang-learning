package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	var a, b string
	a = "hello"
	b = "world"
	a, b = swap(a, b)
	fmt.Println(a, b)
}
