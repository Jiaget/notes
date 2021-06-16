package main

import "fmt"

func main() {
	a := new([]int)
	b := make([]int, 1, 3)
	fmt.Println(a, cap(b))
}
