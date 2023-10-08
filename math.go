package main

import "fmt"

func main() {
	fmt.Println(soma(10, 10))
}

func soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func multiple(a int, b int) int {
	return a * b
}

func divided(a int, b int) int {
	return a / b
}

func sum_and_multiple(a int, b int, c int) int {
	return a + b*c
}
