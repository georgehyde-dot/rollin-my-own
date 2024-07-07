package primal

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to implement the Miller-Rabin Primality Test
func isPrime(n int, k int) bool {
	fmt.Printf("Testing %d for primality\n", n)

	if n == 2 || n == 3 {
		return true
	}

	if n <= 1 || n%2 == 0 {
		return false
	}

	// Find r and d such that n = 2^r * d + 1
	r, d := 0, n-1
	for d%2 == 0 {
		r++
		d /= 2
	}

	// Repeat k times
	for i := 0; i < k; i++ {
		a := rand.Intn(n-2) + 2
		x := powMod(a, d, n)

		if x == 1 || x == n-1 {
			continue
		}

		j := 0
		for ; j < r; j++ {
			x = powMod(x, 2, n)
			if x == n-1 {
				break
			}
		}

		if j == r {
			return false
		}
	}

	return true
}

func powMod(x, y, m int) int {
	r := 1
	x = x % m
	for y > 0 {
		if y%2 == 1 {
			r = (r * x) % m
		}
		y = y / 2
		x = (x * x) % m
	}

	return r
}

// A function that uses isPrime to search for a specified amount of time to find a prime number
func FindPrime(maxTimeSeconds int) (int, func()) {
	done := make(chan bool)
	ans := make(chan int)

	cancel := func() {
		close(done)
		close(ans)
	}

	go func() {
		time.Sleep(time.Duration(maxTimeSeconds) * time.Second)
		fmt.Println("Timer Closed")
		done <- true
	}()
	go func() {
		for {
			check := rand.Int()
			if isPrime(check, 16) {
				ans <- check
			}
		}
	}()
	// use to test with known primes
	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	ans <- 17
	// }()
	for {
		select {
		case <-done:
			return 0, cancel
		case x := <-ans:
			return x, cancel
		}
	}
}
