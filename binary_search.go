package main

import "fmt"

func BinarySearch(datalist []int, key int) int {
	left := 0
	right := len(datalist) - 1
	for left <= right {
		mid := (left + right) / 2
		if datalist[mid] == key {
			return mid
		} else if datalist[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	datalist := []int{10, 5, 7, 8, 3, 2, 1}
	fmt.Println(BinarySearch(datalist, 7))
	fmt.Println(BSearch(datalist, 7))
}

//---------------------------------------------------------------------

func BSearch(datalist []int, key int) bool {
	var left = 0
	var right = len(datalist) - 1
	for left <= right {
		var mid = int(left+right) / 2
		var mid_value = datalist[mid]
		if mid_value == key {
			return true
		} else if datalist[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
