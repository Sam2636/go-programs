package main

import "fmt"

func main() {
    // Declare a slice of integers
    numbers := []int{1, 2, 3, 4, 5}

    // Add an element to the end of the slice
    numbers = append(numbers, 6)

    // Insert an element at a specific index
    index := 2
    numbers = append(numbers[:index], append([]int{10}, numbers[index:]...)...)

    // Remove an element at a specific index
    index = 3
    numbers = append(numbers[:index], numbers[index+1:]...)

    // Print the contents of the slice
    fmt.Println(numbers)
}
