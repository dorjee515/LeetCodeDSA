/*
**Kadane's Algorithm - O(n) Time and O(1) Space

The idea of Kadane's algorithm is to traverse over the array from left to right and for each element,
find the maximum sum among all subarrays ending at that element. The result will be the maximum of all these values.
To calculate the maximum sum of subarray ending at current element, say maxEnding, we can use the maximum sum ending at the previous element.
So for any element, we have two choices:
Choice 1: Extend the maximum sum subarray ending at the previous element by adding the current element to it.

	If the maximum subarray sum ending at the previous index is positive, then it is always better to extend the subarray.

Choice 2: Start a new subarray starting from the current element. If the maximum subarray sum ending at the previous index is negative,

	it is always better to start a new subarray from the current element.

This means that maxEnding at index i = max(maxEnding at index (i - 1) + arr[i], arr[i]) and the maximum value of maxEnding at any index
will be our answer.
arr: [2, 3, -8, 7, -1, 2, 3]
*/
package main

func KadaneMaxSum(arr []int) int {
	currMaxSum := arr[0]
	MaxSum := arr[0]

	for i := 1; i < len(arr); i++ {
		currMaxSum = max(arr[i], currMaxSum+arr[i]) //5 -3 7 6 8 11
		MaxSum = max(currMaxSum, MaxSum)            //5 5 7 7 8 11
	}
	return MaxSum
}
