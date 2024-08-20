package main

import (
	"fmt"
)

// Function to sum a slice of numbers and send the result to a channel
func sumPart(nums []int, ch chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	ch <- sum // Send the sum to the channel
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create channels for the two parts
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Split the list and sum each part in parallel
	go sumPart(nums[:5], ch1)
	go sumPart(nums[5:], ch2)

	// Receive the sums from the channels and add them together
	sum1 := <-ch1
	sum2 := <-ch2

	totalSum := sum1 + sum2
	fmt.Println("Total Sum:", totalSum)
}
