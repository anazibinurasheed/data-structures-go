package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 7, 8, 9}

	result := BinarySearch(8, slice...)
	fmt.Println(result)

	left, right := 0, len(slice)-1
	result = RecursiveBinarySearch(left, right, 3, slice...)
	fmt.Println(result)
}

// Binary search is a searching algorithm which is used to search an element index in the array, which is only applicable to the sorted arrays. And it will only take o(logn) complexity.

// th e
func BinarySearch(val int, slice ...int) int {

	left, right := 0, len(slice)-1

	for left <= right {
		//if left and right are both int32 values, and their values are 2147483647 (the maximum value of int32), then the addition operation will overflow, and the value of mid will be incorrect.

		//the expression left + (right - left) / 2 cannot overflow. This is because the subtraction operation is performed before the addition operation, and the result of the subtraction operation is always less than or equal to the value of right.
		mid := left + (right-left)/2

		if slice[mid] == val {
			return mid

		}

		if slice[mid] > val {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return -1
}

func RecursiveBinarySearch(left, right, val int, slice ...int) int {

	if left <= right {
		mid := left + (right-left)/2

		if slice[mid] == val {
			return mid
		}

		if slice[mid] > val {
			return RecursiveBinarySearch(left, mid-1, val, slice...)
		} else {
			return RecursiveBinarySearch(mid+1, right, val, slice...)

		}

	}
	return -1

}
