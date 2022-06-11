/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sat Jun 11 2022
 */

// REFERENCE: https://leetcode.com/problems/n-th-tribonacci-number/

// package main

// import (
// 	"fmt"
// )

func recursiveTribonacci(n int, tribonacciResult []int) int {
	if n == 0 {
		return 0
	}
	if tribonacciResult[n] != 0 {
		return tribonacciResult[n]
	}
	tribonacciResult[n] = recursiveTribonacci(
		n-1,
		tribonacciResult) + recursiveTribonacci(
		n-2,
		tribonacciResult) + recursiveTribonacci(
		n-3,
		tribonacciResult)

	return tribonacciResult[n]
}

func tribonacci(n int) int {
	tribonacciResult := make([]int, 38)
	tribonacciResult[1] = 1
	tribonacciResult[2] = 1

	return recursiveTribonacci(n, tribonacciResult)
}

// func main() {
// 	var n int
// 	for {
// 		fmt.Scanln(&n)

// 		if n < 0 {
// 			break
// 		}

// 		fmt.Println("output : ", tribonacci(n))
// 	}
// }

// Test cases

// 4
// 4
// 8
// 44
// 7
// 24
// 25
// 1389537
// 10
// 149
// 12
// 504
// 37
// 2082876103
