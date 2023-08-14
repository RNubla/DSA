package main

import "fmt"

func min(a, b *int) int {
	if *(a) < *(b) {
		return *(a)
	}
	return *(b)
}

func max(a, b *int) int {
	if *(a) > *(b) {
		return *(a)
	}
	return *(b)
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := 0
	for left != right {
		minHeight := min(&height[left], &height[right])
		area := minHeight * (right - left)
		result = max(&result, &area)
		if height[left] <= height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return result
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
