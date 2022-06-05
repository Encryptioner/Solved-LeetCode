/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sun Jun 05 2022
 */

// package main

// import (
// 	"fmt"
// 	"sort"
// )

func minMoves2(nums []int) int {
	sort.Ints(nums)

	fmt.Println(nums)

	i := 0
	j := len(nums) - 1
	minMoves := 0

	for {
		if i >= j {
			break
		}
		minMoves += (nums[j] - nums[i])
		i++
		j--
	}
	return minMoves
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
// 		fmt.Println("output : ", minMoves2(nums))
// 	}
// }

// Test cases

// [1,2,3]
// 2
// [1,10,2,9]
// 16
// [-1, 5, 1]
// 6
// [0,-1,-1,5,3]
// 10
// [5,0,2,-1,5,5]
// 14
// [1,0,0,8,6]
// 14
// [203125577,-349566234,230332704,48321315,66379082,386516853,50986744,-250908656,-425653504,-212123143]
// 2127271182
// [1]
// 0
// [1,1,2]
// 1
