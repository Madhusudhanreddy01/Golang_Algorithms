package main

import "fmt"

// func LinearSearch(datalist []int, key int) int {
// 	a := 0
// 	for a < len(datalist) {
// 		if datalist[a] == key {
// 			return a
// 		}
// 		a++
// 	}
// 	return -1
// }

func main() {
	items := []int{10, 5, 7, 8, 3, 2, 1}
	fmt.Println(linearSearch(items, 3))
	// fmt.Println(LSearch(items, 7))
}

// //----------------------------------------------------------------------------

// func LSearch(datalist []int, key int) bool {
// 	for _, item := range datalist {
// 		if item == key {
// 			return true
// 		}
// 	}
// 	return false
// }

// // --------------------------------------------------------------------------------

// package main

// import "fmt"

func linearSearch(A []int, data int) int {
	for i, item := range A {
		if item == data {
			return i + 1
		}
	}
	return -1
}

// func main() {
// 	A := []int{67, 68, 16, 8, 5, 86, 29, 21, 50}
// 	a := linearSearch(A, 5)
// 	fmt.Printf("The Source Array : %v\n", A)
// 	fmt.Printf("The element %v is found at %v location", 5, a)
// }
