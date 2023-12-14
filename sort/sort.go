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

func InsertionSort(A []int) []int {
	n := len(A)
	for i := 1; i <= n-1; i++ {
		j := i
		for j > 0 {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
			}
			j -= 1
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

	A = SelectionSort(A)
	fmt.Println("\n After Selection Sorting")
	for _, val := range A {
		fmt.Println(val)
	}

	A = InsertionSort(A)
	fmt.Println("\n After Insertion Sorting")
	for _, val := range A {
		fmt.Println(val)
	}
}
