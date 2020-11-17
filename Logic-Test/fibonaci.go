package main

import "fmt"

func fibonacci (N int) int {
	if N <= 0 || N > 1000 { // limit only between 0 - 1000
		return 0
	}

	if N == 1 { // base case
		return 1
	}
	return fibonacci(N - 1) + fibonacci(N - 2)
}

func main() {
	var N int

	_, _ = fmt.Scanln(&N)
	if N > 0 && N <= 1000 {
		for i := 0; i < N; i++ {
			fmt.Printf("%d ", fibonacci(i))
		}
	}

}
