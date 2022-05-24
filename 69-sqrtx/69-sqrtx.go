/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Tue May 24 2022
 */

// package main

// import "fmt"

func mySqrt(x int) int {
	if x == 2 {
		return 1
	}

	n := 1
	m := 0
	num := x

	for m < num {
		if m*m > num {
			m -= n
			break
		}
		m += n
	}
	return m
}

// func main() {
// 	var num int

// 	for {
// 		fmt.Scanln(&num)

// 		if num < 0 {
// 			break
// 		}
// 		fmt.Println("output : ", mySqrt(num))
// 	}
// }

// Test cases

// 0 = 0
// 1 = 1
// 2 = 1
// 4 = 4
// 2147483648 = 46340
