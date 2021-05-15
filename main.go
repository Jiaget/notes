package main

import (
	"fmt"
	"io"
)

func main() {

	for {
		var n int
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		res := ""
		var backtrace func(x int) (res string)
		backtrace = func(x int) (res string) {
			if x == n {
				return
			}

			x = f1(x)
			res += backtrace(x) + "2"
			res = res[:len(res)-1]

			x = f2(x)
			res += backtrace(x) + "3"
			res = res[:len(res)-1]
			return
		}
		backtrace(0)
		fmt.Println(res)
	}
}

func f1(x int) int {
	return x*2 + 1
}

func f2(x int) int {
	return x*2 + 2
}
