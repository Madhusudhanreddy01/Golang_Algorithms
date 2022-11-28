package main

import "fmt"

func LinearSearch(datalist []int, key int) int {
	a := 0
	for a < len(datalist) {
		if datalist[a] == key {
			return a
		}
		a++
	}
	return -1
}

func main() {
	items := []int{10, 5, 7, 8, 3, 2, 1}
	fmt.Println(LinearSearch(items, 33))
	fmt.Println(LSearch(items, 7))
}

//----------------------------------------------------------------------------

func LSearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}
