/*
*Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/

package main

func productExceptSelf(nums []int) []int {
	pd := 1
	pdWz := 1
	cntOf0 := 0
	dupNums := make([]int, 0)
	for _, num := range nums {
		if num == 0 {
			cntOf0++
		}
		pd *= num
		dupNums = append(dupNums, num)
	}
	//-1,1,0,-3,3
	for i, num := range nums {
		if num != 0 {
			nums[i] = pd / num
		} else {
			for j := 0; j < len(dupNums); j++ {
				if j != i {
					pdWz *= dupNums[j]
				}
			}
			nums[i] = pdWz
		}
	}
	return nums
}
