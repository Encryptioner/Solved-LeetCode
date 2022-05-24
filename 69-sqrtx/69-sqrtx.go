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
	// works for 1 decimal accuracy after decimal
	n := 1
	// n := 0.1

	m := float64(0)
	num := float64(x)

	for m < num {
		// NOTE: floating point arithmetic is not reliable, for example 0.7+0.1 becomes 0.799999999
		// So, multiple with pow(10, decimal accuracy after decimal)
		if m*m > num {
			m = ((m * 10) - float64(n)) / 10
			// m -= n
			break
		}
		m = ((m * 10) + float64(n)) / 10
		// m += n
	}
	return int(m)
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
// 2 = 2
// 4 = 4
// 2147483648 = 46340
