package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("这是一个 《冒泡排序，简单快排， 带二分思想的快排》 的排序效率测试")
	var n int
	fmt.Print("请输入待排序数组长度:")
	fmt.Scanf("%d", &n)
	testNums := generate(n, 0, 1000)
	testquickSort1(testNums)
	testquickSort2(testNums)
	testbubble(testNums)

}

func testquickSort1(nums []int) {
	copy := nums[:]
	start := time.Now()
	_ = quickSort1(copy)
	end := time.Now()
	time := start.Sub(end)
	fmt.Printf("简单快排所花费的时间：%v\n", time)
}
func testquickSort2(nums []int) {
	copy := nums[:]
	start := time.Now()
	quickSort2(copy, 0, len(copy)-1)
	end := time.Now()
	time := start.Sub(end)
	fmt.Printf("二分思想的快排所花费的时间：%v\n", time)
}
func testbubble(nums []int) {
	copy := nums[:]
	start := time.Now()
	_ = bubbleSort(copy)
	end := time.Now()
	time := start.Sub(end)
	fmt.Printf("冒泡排序所花费的时间：%v\n", time)
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
		for j := i; j < length-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
