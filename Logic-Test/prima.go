package main

import "fmt"

func prima (n int) string {
	if n < 0 || n > 1000 { // limitation from the question
		return "wrong number"
	}

	if n > 3 {
		if n % 2 == 0 || n % 3 == 0 {
			return "bukan bilangan prima"
		}
	} else if n < 2 {
		return "bukan bilangan prima"
	}

	return "bilangan prima"
}

func main() {
	var n int

	_, _ = fmt.Scanln(&n)
	fmt.Println(prima(n))
}
