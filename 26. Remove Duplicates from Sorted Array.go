package main

import "fmt"

// Remove Duplicates from Sorted Array
// We use two index i and j where j is settled to be i + 1 as initiliazed
func removeDuplicates(nums []int) int {
	i, j, size := 0, 1, len(nums)

	for i = 0; j < size; j++ {
		if nums[i] == nums[j] {
			continue
		}
		i++
		nums[i], nums[j] = nums[j], nums[i]
	}
	return i+1
}

func main(){
	elements := []int{10, 20, 30, 10, 10, 20, 40}
	fmt.Println(removeDuplicates(elements))
}