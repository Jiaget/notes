package main

import (
	"fmt"
)

func main() {
	a := [10]rune{}
	b := make([]rune, 10)
	fmt.Println(len(a), len(b))
}
