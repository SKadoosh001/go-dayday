package main

import "fmt"

func search001(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func search002(arr []int, target int) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 5, 9, 12, 46, 55, 77, 88, 99}
	fmt.Println(a)
	//target := Er_fen_fa(a, 99)
	//println(target)
	target := 1
	println(search001(a, target))
	println(search002(a, target))

	//fmt.Printf("target = %d\n", target)
}
