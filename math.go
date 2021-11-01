package main

import "fmt"

func main() {
	total := Soma(2, 4)
	fmt.Printf("A soma eh: %d\n", total)
}
func Soma(a int, b int) int {
	return a + b
}

func Substracao(a int, b int) int {
	return b - a
}
