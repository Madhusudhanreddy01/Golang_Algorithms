// O(n2)
package main

import "fmt"

func bubbleSort(A []int) []int {
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

func main() {
	A := []int{3, 4, 5, 2, 1}
	fmt.Println(bubbleSort(A))
	// A = bubbleSort(A)
	// fmt.Println("\n After Bubble Sorting")
	// for _, val := range A {
	// 	fmt.Println(val)
	// }
}

// O(n)-Improved Version

// package main
// import "fmt"

// func bubbleSort(A []int) []int {
// 	var sorted bool
// 	items := len(A)
// 	// We run the outer loop until we have sorted all the items.
// 	for !sorted {
// 		// In each iteration we are going to change sorted to true.
// 		sorted = true
// 		// Now we're going to range over our slice checking if they're sorted or not.
// 		for i := 1; i < items; i++ {
// 			// If they're not sorted we swap them and change sorted to false to loop again.
// 			if A[i-1] > A[i] {
// 				A[i-1], A[i] = A[i], A[i-1]
// 				sorted = false
// 			}
// 		}
// 	}
// 	return A
// }

// func main() {
// 	A := []int{3, 4, 5, 2, 1}
// 	A = bubbleSort(A)
// 	fmt.Println("\n After Bubble Sorting")
// 	for _, val := range A {
// 		fmt.Println(val)
// 	}
// }
