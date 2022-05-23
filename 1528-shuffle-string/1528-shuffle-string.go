/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Mon May 23 2022
 */

// REFERENCE: https://leetcode.com/problems/shuffle-string/

// package main

// import (
// 	"fmt"
// )

func restoreString(s string, indices []int) string {

	strMap := map[int]string{}

	for i, v := range indices {
		strMap[v] = string(s[i])
	}

	result := ""
	for i, _ := range indices {
		result += strMap[i]
	}
	return result
}

// func main() {
// 	var str string
// 	for {
// 		fmt.Scanln(&str)

// 		if str == "0" {
// 			break
// 		}
// 		strLength := len(str)
// 		var indices = make([]int, strLength)

// 		for i, _ := range indices {
// 			fmt.Scanf("%d", &indices[i])
// 		}

// 		fmt.Println("output : ", restoreString(str, indices))
// 	}
// }

// Test cases

// s = "acb", indices = [0,2,1], output = abc
// s = "abc", indices = [0,1,2], output = abc
// s = "codeleet", indices = [4,5,6,7,0,2,1,3], output = leetcode
