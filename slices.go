package main

import "fmt"

func main() {
	mySlice := make([]int, 1, 4)
	fmt.Printf("Length is: %d Capacity is: %d", len(mySlice), cap(mySlice))

	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity is: %d", cap(mySlice))
	}

	fmt.Println("")

	for i, _ := range mySlice {
		fmt.Println("For range loop:", i)
	}

	mySlice = []int{1, 2, 3, 4, 5}
	newSlice := []int{10, 20, 30}
	mySlice = append(mySlice, newSlice...)
	fmt.Println("")
	fmt.Println(mySlice)
}
