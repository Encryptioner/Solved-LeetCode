/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sun May 22 2022
 */

package main

import "fmt"

func getEachDigit(digits *[]int, num int) {
	if num >= 10 {
		getEachDigit(digits, num/10)
	}
	*digits = append(*digits, (num % 10))
}

func countEven(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		digits := make([]int, 0)
		getEachDigit(&digits, i)

		sum := 0

		for _, v := range digits {
			sum += v
		}

		if sum %2 == 0 {
			count += 1
		}
	}
	return count
}

// func main() {
// 	var num int
// 	for {
// 		fmt.Scanln(&num)

// 		if num > 1000 {
// 			break
// 		}
// 		fmt.Println("OUTPUT : ", countEven(num))
// 	}
// }
