package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	prev := 0
	total := 0

	for i := len(s) - 1; i >= 0; i-- {
		current := romanMap[s[i]]

		if current < prev {
			total -= current
		} else {
			total += current
		}
		prev = current
	}

	return total
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
