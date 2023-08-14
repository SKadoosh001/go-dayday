package main

import "fmt"

func solution001(arr []int) []int {
	low, high1, high2 := 0, len(arr)-1, len(arr)-1
	ret := make([]int, len(arr))

	for low <= high1 {
		if arr[low]*arr[low] >= arr[high1]*arr[high1] {
			ret[high2] = arr[low] * arr[low]
			low++
		} else {
			ret[high2] = arr[high1] * arr[high1]
			high1--
		}
		high2--
	}
	return ret
}
func main() {
	arr := []int{-4, -3, -2, -1, 0, 1, 2, 9}
	ret := solution001(arr)

	fmt.Println(ret)
}
