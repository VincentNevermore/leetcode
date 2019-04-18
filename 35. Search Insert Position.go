package main

import "fmt"

//the idea of this question is to first check if the target num is less than the first number in the array or not
// if yes return 0 else we move forward
func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}

		if target < nums[i] && target > nums[i-1] {
			return i
		}
	}
	// The program should never reach to this point other wise it is an error program
	return -1
}

func main() {
	nums := []int{1, 3}
	target := 3
	fmt.Println(searchInsert(nums, target))
}
