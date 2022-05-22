/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sun May 22 2022
 */

// REFERENCE: https://leetcode.com/problems/integer-to-roman/

// package main

// import (
// 	"fmt"
// )

// splits the digit of number in array
func getEachDigitReversed(digits *[]int, num int, ch chan bool) {
	for num != 0 {
		*digits = append(*digits, (num % 10))
		num /= 10
	}
	ch <- true
}

func setMappedValuesInPosition(mappedValuesInPosition *[4]map[int]string, ch chan bool) {

	mappedValuesInPosition[0] = map[int]string{
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		8: "VIII",
		9: "IX",
	}

	mappedValuesInPosition[1] = map[int]string{
		1: "X",
		2: "XX",
		3: "XXX",
		4: "XL",
		5: "L",
		6: "LX",
		7: "LXX",
		8: "LXXX",
		9: "XC",
	}

	mappedValuesInPosition[2] = map[int]string{
		1: "C",
		2: "CC",
		3: "CCC",
		4: "CD",
		5: "D",
		6: "DC",
		7: "DCC",
		8: "DCCC",
		9: "CM",
	}

	mappedValuesInPosition[3] = map[int]string{
		1: "M",
		2: "MM",
		3: "MMM",
	}

	ch <- true
}

func intToRoman(num int) string {

	digits := make([]int, 0)
	digitsCh := make(chan bool)

	var mappedValuesInPosition = [4]map[int]string{}
	mappedValuesInPositionCh := make(chan bool)

	go getEachDigitReversed(&digits, num, digitsCh)
	go setMappedValuesInPosition(&mappedValuesInPosition, mappedValuesInPositionCh)

	<-digitsCh
	<-mappedValuesInPositionCh

	value := ""
	for i := len(digits) - 1; i >= 0; i-- {
		val := mappedValuesInPosition[i][digits[i]]
		value += val
	}

	return value
}

// func main() {
// 	var num int

// 	for {
// 		fmt.Scanln(&num)

// 		if num > 3999 {
// 			break
// 		}
// 		fmt.Println("ROMAN : ", intToRoman(num))
// 	}
// }

// Test cases

// CM = 900
// DXVI =  516
// MCLXXXIII =  1183
// MDCCXXXIII =  1733
// MDCCCXLI =  1841
// DCCCLXXI =  871
// MCCCLXXXIII =  1383
// MCCL =  1250
// MDLXVIII =  1568
// MMCXIII =  2113
// MMCCLXXIV =  2274
// MCMLXXXII = 1977
// CMXCIX = 999
// MCMLXXXVIII = 1988
