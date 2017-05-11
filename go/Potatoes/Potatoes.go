package main

import "fmt"

func Potatoes(p0, w0, p1 int) int {
	return w0 * (100 - p0) / (100 - p1)
}

func main() {
	fmt.Println(Potatoes(99, 100, 98))
	fmt.Println(Potatoes(82, 127, 80))
}
