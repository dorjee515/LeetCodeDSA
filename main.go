package main

import "fmt"

func main() {
	//TwoSum
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(nums, target)
	fmt.Println(ans)
	//Best Time to Buy and Sell Stock
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	//kande's algo
	arr := []int{2, 3, -8, 7, -1, 2, 3}
	fmt.Println(KadaneMaxSum(arr))
	//Contains Duplicate
	duplicates := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(duplicates))
	//ProductOfArrayExceptSelf
	nums1 := []int{1, 2, 3, 4}
	ans1 := productExceptSelf(nums1)
	for _, num := range ans1 {
		fmt.Println(num)
	}
	//MaximumSubarray
	nums2 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums2))
}
