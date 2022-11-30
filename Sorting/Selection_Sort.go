package main

import "fmt"

func selectionSort(A []int) []int {
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
	A := []int{3, 4, 5, 2, 1}
	fmt.Println(selectionSort(A))
	// A = selectionSort(A)
	// fmt.Println("\n After Selection Sorting")
	// for _, val := range A {
	// 	fmt.Println(val)
	// }
}
