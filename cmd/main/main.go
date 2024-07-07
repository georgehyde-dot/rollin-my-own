package main

import (
	"fmt"
	"os"

	"github.com/georgehyde-dot/rollin-my-own/pkg/primal"
)

func main() {
	fmt.Println("Do Something")
	num, cancel := primal.FindPrime(3)
	cancel()
	if num != 0 {
		fmt.Printf("Prime! %d\n", num)
		os.Exit(1)
	}
	fmt.Println("No prime")
	os.Exit(0)

}
