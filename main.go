package main

import (
	"fmt"
)

// Function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Goroutine to generate numbers and send them to a channel
func generateNumbers(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch) // Close the channel after sending all numbers
}

// Goroutine to filter primes and send them to another channel
func filterPrimes(ch chan int, primeCh chan int) {
	for num := range ch {
		if isPrime(num) {
			primeCh <- num
		}
	}
	close(primeCh) // Close the channel after sending all primes
}

func main() {
	numberCh := make(chan int)
	primeCh := make(chan int)

	// Start the goroutines
	go generateNumbers(numberCh)
	go filterPrimes(numberCh, primeCh)

	// Print the prime numbers
	for prime := range primeCh {
		fmt.Println(prime)
	}
}
