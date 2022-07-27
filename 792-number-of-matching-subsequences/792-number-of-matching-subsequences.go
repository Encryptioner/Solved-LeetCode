/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Wed Jul 27 2022
 */

// REFERENCE: https://leetcode.com/problems/number-of-matching-subsequences/

// package main

// import (
// 	"fmt"
// 	"os"
// )

func numMatchingSubseq(s string, words []string) int {
	initialCnt := 0
	cnt := 0
	strLen := len(s)
	wordLength := len(words)
	wordIndexes := make([]int, wordLength)

	for strIndex := 0; strIndex < strLen; strIndex++ {

		initialCnt = 0

		for i := 0; i < wordLength; i++ {
			word := words[i]
			currentIndex := wordIndexes[i]

			if currentIndex >= len(word) {
				initialCnt++
				continue
			}

			if word[currentIndex] == s[strIndex] {
				wordIndexes[i] = currentIndex + 1
			}
		}

		if initialCnt == wordLength {
			return wordLength
		}
	}

	for i := 0; i < wordLength; i++ {
		word := words[i]
		currentIndex := wordIndexes[i]

		if currentIndex == len(word) {
			cnt++
		}
	}

	return cnt
}

// func main() {
// 	var s string
// 	var wordLength int

// 	dir, err := os.Getwd()
// 	if err != nil {
// 		panic(err)
// 	}

// 	inputFile := dir + "/leetcode/792/input.txt"
// 	outputFile := dir + "/leetcode/792/output.txt"

// 	os.Stdin, err = os.OpenFile(inputFile,
// 		os.O_RDONLY|os.O_CREATE, 0666)

// 	os.Stdout, err = os.OpenFile(outputFile,
// 		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

// 	for {
// 		fmt.Scanln(&wordLength)
// 		if wordLength <= 0 {
// 			break
// 		}
// 		fmt.Scanln(&s)

// 		words := make([]string, wordLength)

// 		for i := 0; i < wordLength; i++ {
// 			fmt.Scanf("%s", &words[i])
// 		}

// 		fmt.Println("output : ", numMatchingSubseq(s, words))
// 	}
// }
