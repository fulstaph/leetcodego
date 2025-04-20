package leetcodego

import "math"

// https://leetcode.com/problems/increasing-triplet-subsequence
// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k].
// If no such indices exists, return false.
// Example 1:
// Input: nums = [1,2,3,4,5]
// Output: true
// Explanation: Any triplet where i < j < k is valid.
// Example 2:
// Input: nums = [5,4,3,2,1]
// Output: false
// Explanation: No triplet exists.
// Example 3:
// Input: nums = [2,1,5,0,4,6]
// Output: true
// Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
// Constraints:
// 1 <= nums.length <= 5 * 105
// -231 <= nums[i] <= 231 - 1
// Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?
// Solution:
// We can use two variables to keep track of the smallest and second smallest elements in the array.
// If we find a number that is greater than the second smallest element, we have found an increasing triplet.
// If we find a number that is smaller than the smallest element, we update the smallest element.
// If we find a number that is between the smallest and second smallest elements, we update the second smallest element.
// If we reach the end of the array without finding an increasing triplet, we return false.
// Time complexity: O(n)
// Space complexity: O(1)
func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}
	return false
}
