/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Wed Jul 20 2022
 */

// REFERENCE: https://leetcode.com/problems/number-of-matching-subsequences/

// package main

// import (
// 	"fmt"
// 	"os"
// )

/* Returns true if s1 is subsequence of s2*/
func isSubsequence(s1 string, s2 string) bool {

	n := len(s1)
	m := len(s2)
	i := 0
	j := 0

	for i < n && j < m {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}

	/*
		If i reaches end of s1,that mean we found all
		characters of s1 in s2,
		so s1 is subsequence of s2, else not
	*/
	return i == n
}

func numMatchingSubseq(s string, words []string) int {
	cnt := 0
	for i := 0; i < len(words); i++ {
		if isSubsequence(words[i], s) {
			cnt += 1
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
