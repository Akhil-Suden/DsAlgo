package main

import "fmt"

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

func main() {
	A := []int{3, 4, 5, 2, 1}
	A = BubbleSort(A)
	fmt.Println("\n After Bubble Sorting")
	for _, val := range A {
		fmt.Println(val)
	}
}
