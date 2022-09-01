package main

import "fmt"

func InsertionAtBeginning(data int) {
	if len(arr) == length {
		fmt.Println("The array is overflowed!")
		return
	}

	for i := len(arr) - 1; i >= 0; i-- {
		arr[i+1] = arr[i]
	}
	arr[0] = data
}

func InsertionAtEnd(data int) {
	if len(arr) == length {
		fmt.Println("The array is overflowed!")
		return
	}

	arr[len(arr)-1] = data
}

func InsertionAtPosition(data, pos int) {
	if len(arr) == length {
		fmt.Println("The array is overflowed!")
		return
	}

	for i := len(arr) - 1; 
}