package main

import "fmt"

func SelectionSort(A []int) []int {
	var n = len(A)
	for i := 0; i < n; i++ {
		var minIndex = i
		for j := i; j < n; j++ {
			if A[j] < A[minIndex] {
				minIndex = j
			}
		}
		A[i], A[minIndex] = A[minIndex], A[i]
	}
	return A
}

func main() {
	arr := []int{10, 4, 43, 5, 57, 91, 45, 9, 7}
	SelectionSort(arr)
	fmt.Print(arr)

}
