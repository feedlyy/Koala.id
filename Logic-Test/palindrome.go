package main

import "fmt"

func palindromeCheck (s string) bool {
	if len(s) < 2 { // validation if string is too short
		return false
	}

	if s[0:1] == "-" { // remove "-" on negatives numbers
		s = s[1:]
	}

	a := 0
	b := len(s) - 1

	for b > 1 {
		if string(s[a]) != string(s[b]) { //check string from left and right
			return false
		}
		a++
		b--
	}
	return true
}

func main() {
	var s string
	_, _ = fmt.Scanln(&s)

	fmt.Println(palindromeCheck(s))
}
