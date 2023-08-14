package main

import "fmt"

func twoSum(list []int, target int) []int {
	cache := make(map[int]int)

	for i, v := range list {
		diff := target - v
		val, exist := cache[diff]
		// fmt.Println(val)
		if exist {
			return []int{i, val}
		}
		cache[v] = i
	}
	// fmt.Println(cache)
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
