package algorithm

import (
	"fmt"
	"math/rand"
	"time"
)

// 排序算法测试
func SortTest() {
	fmt.Println("这是一个 【冒泡排序，简单快排， 优化快排， 归并排序】 的排序效率测试")
	var n int
	fmt.Print("请输入待排序数组长度:")
	fmt.Scanf("%d", &n)
	testNums := Generate(n, 0, 1000)
	TestquickSort1(testNums)
	TestquickSort2(testNums)
	Testbubble(testNums)
	TestMergeSort(testNums)
}

func TestMergeSort(nums []int) {
	copied := make([]int, len(nums))
	copy(copied, nums)
	start := time.Now()
	_ = mergeSort(copied)
	end := time.Now()
	time := end.Sub(start)
	fmt.Printf("归并排序所花费的时间：%v\n", time)
}

func TestquickSort1(nums []int) {
	copied := make([]int, len(nums))
	copy(copied, nums)
	start := time.Now()
	_ = quickSort1(copied)
	end := time.Now()
	time := end.Sub(start)
	fmt.Printf("简单快排所花费的时间：%v\n", time)
}
func TestquickSort2(nums []int) {
	copied := make([]int, len(nums))
	copy(copied, nums)
	start := time.Now()
	quickSort2(copied, 0, len(copied)-1)
	end := time.Now()
	time := end.Sub(start)
	fmt.Printf("优化快排所花费的时间：%v\n", time)
}
func Testbubble(nums []int) {
	copied := make([]int, len(nums))
	copy(copied, nums)
	start := time.Now()
	_ = bubbleSort(copied)
	end := time.Now()
	time := end.Sub(start)
	fmt.Printf("冒泡排序所花费的时间：%v\n", time)
}

func Generate(size, min, max int) []int {
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

func mergeSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	mid := length / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

// left, right 是两个有序数组，我们要将它们有序地合并形成新的有序数组
func merge(left []int, right []int) (res []int) {
	left_size := len(left)
	right_size := len(right)
	i, j := 0, 0

	for {
		// 如果left 先遍历完, right[j:]未遍历的数肯定比left[i]大，合并在右边就行
		if i == left_size {
			res = append(res, right[j:]...)
			break
		}
		if j == right_size {
			res = append(res, left[i:]...)
			break
		}

		if left[i] > right[j] {
			res = append(res, right[j])
			j++
		} else {
			res = append(res, left[i])
			i++
		}
	}
	return
}
