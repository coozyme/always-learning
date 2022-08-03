package main

import "fmt"

func main() {
	ua := []int{2, 7, 11, 15}
	twoSum(ua, 9)
}

func twoSum(nums []int, target int) []int {
	fmt.Println(nums)
	for i := 0; i <= len(nums); i++ {
		for j := 0; j < len(nums)-j+1; j++ {

			if nums[i]+nums[i]+1 == 9 {
				fmt.Println(nums[i] + nums[i] + 1)
				nums = append(nums, nums[i])
				nums = append(nums, nums[i]+1)
				// return nums
			}
		}
	}
	return nums
}
