/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Wed Jun 01 2022
 */

// REFERENCE: https://leetcode.com/problems/count-operations-to-obtain-zero/

// package main

// import "fmt"

func countOperations(num1 int, num2 int) int {
	operationCount := 0

	if (num1 == 0) || (num2 == 0) {
		return operationCount
	}

	for {
		operationCount += 1

		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}

		if (num1 == 0) || (num2 == 0) {
			break
		}
	}

	return operationCount
}

// func main() {
// 	var num1 int
// 	var num2 int
// 	for {
// 		fmt.Scanln(&num1)
// 		fmt.Scanln(&num2)
// 		if (num1 == -1) || (num2 == -1) {
// 			break
// 		}
// 		fmt.Println("output : ", countOperations(num1, num2))
// 	}
// }

// Test cases

// num1 = 2, num2 = 3, output = 6
// num1 = 10, num2 = 10, output = 1
// num1 = 10, num2 = 12, output = 6
// num1 = 0, num2 = 10000, output = 0
// num1 = 1, num2 = 100000, output = 100000
// num1 = 99, num2 = 100, output = 100
// num1 = 100, num2 = 99, output = 100
// num1 = 101, num2 = 58921, output = 600
