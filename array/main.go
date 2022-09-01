package main

import "fmt"

const (
	length int = 10
)

var (
	arr [length]int
)

func main() {

	var choice int

	for {
		fmt.Println("Please choose an option from below : ")
		fmt.Println("1. Insert at beginning\n2. Insert at end\n3. Insert at position\n4. Delete from start\n5. Delete from end\n6. Delete from position\n7. Display the elements\n8. Exit the program")
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			var data int
			fmt.Println("Enter the data to insert : ")
			fmt.Scanf("%d\n", &data)
			InsertionAtBeginning(data)
		case 2:
			var data int
			fmt.Println("Enter the data to insert : ")
			fmt.Scanf("%d\n", &data)
			InsertionAtEnd(data)
		case 3:
			var data, pos int
			fmt.Println("Enter the data to insert : ")
			fmt.Scanf("%d\n", &data)
			fmt.Println("Enter the position to insert the data : ")
			fmt.Scanf("%d\n", &pos)
			InsertionAtPosition(data, pos)
		}
	}

}
