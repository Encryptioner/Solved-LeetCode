/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Wed May 25 2022
 */

// REFERENCE: https://leetcode.com/problems/jump-game/

// package main

// import "fmt"

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	lastIndex := len(nums) - 1
	nearestGoodIndex := lastIndex

	for i := lastIndex; i > 0; i-- {
		checkingIndex := nums[i] + i
		if checkingIndex >= nearestGoodIndex {
			nearestGoodIndex = i
		}
	}
	firstIndexValue := nums[0]
	return firstIndexValue >= nearestGoodIndex
}

// func main() {
// 	var num int
// 	for {
// 		fmt.Scanln(&num)
// 		if num == -1 {
// 			break
// 		}

// 		var nums = make([]int, num)

// 		for i, _ := range nums {
// 			fmt.Scanf("%d", &nums[i])
// 		}
// 		fmt.Println("output : ", canJump(nums))
// 	}
// }

// Test cases

// [2,3,1,1,4] = true
// [3,2,1,0,4] = false
// [1] = true
// [2,0] = true
