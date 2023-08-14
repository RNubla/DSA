package main

import (
	"fmt"
	"sort"
)

/*
Reason for Sorting
The reason why we sort the input array of strings and compare the first and last strings is that the longest common prefix of all the strings must be a prefix of the first string and a prefix of the last string in the sorted array. This is because strings are ordered based on their alphabetical order (Lexicographical order).
For example, consider the input array of strings {"flower", "flow", "flight"}. After sorting the array, we get {"flight", "flow", "flower"}. The longest common prefix of all the strings is "fl", which is located at the beginning of the first string "flight" and the second string "flow". Therefore, by comparing the first and last strings of the sorted array, we can easily find the longest common prefix.
*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestCommonPrefix(strs []string) string {
	var result []byte
	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]

	for i := 0; i < min(len(first), len(last)); i++ {
		fmt.Println("first[i]:", string(first[i]))
		fmt.Println("last[i]:", string(last[i]))
		if first[i] != last[i] {
			return string(result)
		}
		result = append(result, first[i])
		// fmt.Println(string(result))
	}

	return string(result)
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
