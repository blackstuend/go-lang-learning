package main

import "fmt"

func Fibonacci(num int) int {
	array := make([]int, 100)

	array[0] = 0
	array[1] = 1

	for i := 2; i <= num; i++ {
		array[i] = array[i-1] + array[i-2]
	}
	return array[num]
}

func main() {
	result := Fibonacci(5)

	fmt.Printf("fibonacci result=%+v\n", result)
}
