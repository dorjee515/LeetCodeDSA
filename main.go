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
}
