package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// sort the input array
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// skip dupes
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, 1, 1}

	result := threeSum(nums)
	fmt.Println(result)
	// for _, triplet := range result {
	// 	fmt.Println(triplet)
	// }
}

/*
Debugging session
nums = -4, -1,-1, 0, 1, 2

i = 0
    left = 0 + 1 = 1
    right = 6 - 1 = 5
    while left < right //1st
        sum = -4 + -1 + 2 = -3
        left = 2
        right = 5
    while left < right //2nd
        sum = -4 + -1 + 2 = -3
        left = 3
        right = 5
    while left < right //3rd
        sum = -4 + 0 + 2 = -2
        left = 4
        right = 5
    while left < right //4
        sum = -4 + 1 + 2 = -1
        left = 5
        right = 5

result = [-1, -1, 2]
i = 1
    left = 1 + 1 = 2
    right = 6 - 1 = 5
    while left < right //1st
        sum = -1 + -1 + 2 = 0
        left = 3
        right = 4

    while left < right //2nd
        sum = -1 + 0 + 1 = 0
        left = 4
        right = 3



*/
