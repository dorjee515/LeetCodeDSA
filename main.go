package main

import "fmt"

func main() {
	//TwoSum
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("Two Sum")
	ans := twoSum(nums, target)
	fmt.Println(ans)
	//Best Time to Buy and Sell Stock
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Max Profit")
	fmt.Println(maxProfit(prices))
	//kande's algo
	arr := []int{2, 3, -8, 7, -1, 2, 3}
	fmt.Println("kande's algo")
	fmt.Println(KadaneMaxSum(arr))
	//Contains Duplicate
	duplicates := []int{1, 2, 3, 4}
	fmt.Println("Contains Duplicate")
	fmt.Println(containsDuplicate(duplicates))
	//ProductOfArrayExceptSelf
	nums1 := []int{1, 2, 3, 4}
	fmt.Println("Product of Array Except Self")
	ans1 := productExceptSelf(nums1)
	for _, num := range ans1 {
		fmt.Println(num)
	}
	//MaximumSubarray sum
	fmt.Println("MaximumSubarray Sum")
	nums2 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums2))
	//MaximuSubarray Product
	fmt.Println("MaximumSubarray Product")
	nums3 := []int{-2}
	fmt.Println(maxProduct(nums3))
}
