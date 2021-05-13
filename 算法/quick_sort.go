package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start1 := time.Now()
	testNums := generate(10000, 0, 1000)
	sorted := quickSort1(testNums)
	end1 := time.Now()
	time1 := end1.Sub(start1)

	start3 := time.Now()
	sorted = bubbleSort(testNums)
	end3 := time.Now()
	fmt.Println("数组长度", len(sorted))
	// fmt.Printf("待快排数组:%v\n", testNums)
	// fmt.Printf("排序后：%v\n", sorted)
	fmt.Println("------------------")
	fmt.Printf("快排1的用时: %v\n", time1)
	start2 := time.Now()
	quickSort2(testNums, 0, len(testNums)-1)
	end2 := time.Now()
	fmt.Println("快排2的用时：", end2.Sub(start2))
	fmt.Println("冒泡排序：", end3.Sub(start3))

}

func generate(size, min, max int) []int {
	nums := make([]int, size)
	for i := range nums {
		nums[i] = random(0, 10000)
	}
	return nums
}

func random(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func quickSort1(s []int) []int {
	if len(s) < 2 {
		return s
	}

	v := s[0]
	var left, right []int
	for _, e := range s[1:] {
		if e <= v {
			left = append(left, e)
		} else {
			right = append(right, e)
		}
	}

	// 实现了“quickSort(left) + v + quickSort(right)”的操作
	return append(append(quickSort1(left), v), quickSort1(right)...)
}

func quickSort2(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort2(arr, start, j)
		}
		if end > i {
			quickSort2(arr, i, end)
		}
	}
}

func bubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
