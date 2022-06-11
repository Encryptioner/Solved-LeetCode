/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sat Jun 11 2022
 */

// REFERENCE: https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/

// package main

// import (
// 	"fmt"
// 	"sort"
// )

func largestSumAfterKNegations(nums []int, k int) int {
	n := len(nums)

	if n == 1 {
		if k%2 == 1 {
			return nums[0] * (-1)
		}
		return nums[0]
	}

	sort.Ints(nums)

	for i := 0; i < k; i++ {
		value := nums[0] * (-1)
		nums[0] = value

		if value == 0 {
			break
		}

		for j := 1; j < n; j++ {
			jValue := nums[j]
			if jValue >= value {
				break
			}
			nums[j-1] = jValue
			nums[j] = value
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	return sum
}

// func main() {
// 	var n, k int

// 	for {
// 		fmt.Scanln(&n)

// 		if n <= 0 {
// 			break
// 		}
// 		nums := make([]int, n)
// 		for i := 0; i < n; i++ {
// 			fmt.Scanf("%d", &nums[i])
// 		}
// 		fmt.Scanln(&k)

// 		fmt.Println("output : ", largestSumAfterKNegations(nums, k))
// 	}
// }

// Test cases

// [4,2,3]
// 1
// 5
// [3,-1,0,2]
// 3
// 6
// [2,-3,-1,5,-4]
// 2
// 13
// [2,-3,-4,8,-5,-8,0,9,-100]
// 100
// 139
// [99]
// 99
// -99
// [100]
// 100
// 100
// [0]
// 5
// 0
