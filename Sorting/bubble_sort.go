package main

import "fmt"

func main() {
	A := []int{3, 4, 5, 2, 1}
	fmt.Println(BubbleSort(A))
	fmt.Println(BSort(A))
	// fmt.Println("\n After Bubble sorting")
	// for _, val := range A {
	// 	fmt.Println(val)
	// }
}

// O(n2)
func BubbleSort(A []int) []int {
	n := len(A)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}

		}
	}
	return A
}

// -------------------------------------------------------------------

// O(n)
func BSort(A []int) []int {
	var sorted bool
	items := len(A)
	for !sorted {
		sorted = true
		for i := 1; i < items; i++ {
			if A[i-1] > A[i] {
				A[i-1], A[i] = A[i], A[i-1]
				sorted = false
			}
		}
	}
	return A
}
