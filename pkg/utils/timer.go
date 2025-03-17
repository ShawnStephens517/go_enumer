package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Declare a local random generator seeded with the current time.
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateFirstNPrimes returns a slice containing the first n prime numbers.
func GenerateFirstNPrimes(n int) []int {
	primes := []int{}
	num := 2
	for len(primes) < n {
		if IsPrime(num) {
			primes = append(primes, num)
		}
		num++
	}
	return primes
}

// IsPrime checks if a number is prime.
func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	sqrtNum := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrtNum; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// RandomDelay computes the first 500 prime numbers, calculates the square root of (prime * 2)
// for each, and returns a randomly chosen value as a time.Duration.
func RandomDelay() time.Duration {
	primes := GenerateFirstNPrimes(500)
	delays := make([]float64, len(primes))
	for i, p := range primes {
		delays[i] = math.Sqrt(2 * float64(p))
	}
	// Choose a random delay from the computed values.
	randIndex := r.Intn(len(delays))
	return time.Duration(delays[randIndex] * float64(time.Second))
}

// WaitForNextCheck prints the wait time and pauses execution for the randomly computed delay.
func WaitForNextCheck() {
	delay := RandomDelay()
	fmt.Printf("Waiting for %.2f seconds before next check...\n", delay.Seconds())
	time.Sleep(delay)
}
