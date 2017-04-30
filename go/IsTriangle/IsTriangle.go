package main

import "fmt"

func IsTriangle(a, b, c int) bool {

  return a < (b + c) && b < (a + c) && c < (b + a)
}

func main() {
	fmt.Println(IsTriangle(5,1,2)) // false
	fmt.Println(IsTriangle(4, 2, 3)) // true
}
