/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Thu Aug 25 2022
 */

// REFERENCE: https://leetcode.com/problems/ransom-note/

// package main

import (
	"fmt"
	"sort"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteLen := len(ransomNote)
	magazineLen := len(magazine)

	if ransomNoteLen > magazineLen {
		return false
	}

	ransomNoteIndex := 0
	magazineIndex := 0

	ransomNote = SortStringByCharacter(ransomNote)
	magazine = SortStringByCharacter(magazine)

	for ; magazineIndex < magazineLen; magazineIndex++ {

		if ransomNoteIndex >= ransomNoteLen {
			break
		}
		if ransomNote[ransomNoteIndex] == magazine[magazineIndex] {
			ransomNoteIndex++
			continue
		}
	}
	return ransomNoteIndex == ransomNoteLen
}

// func main() {
// 	var ransomNote, magazine string

// 	for {
// 		fmt.Scanln(&ransomNote)

// 		if len(ransomNote) <= 0 {
// 			break
// 		}
// 		fmt.Scanln(&magazine)

// 		fmt.Println("output : ", canConstruct(ransomNote, magazine))
// 	}
// }

// Test cases

// Example launch.json for vscode

// {
//   "version": "0.2.0",
//   "configurations": [
//     {
//       "name": "Launch Package",
//       "type": "go",
//       "request": "launch",
//       "mode": "auto",
//       "console": "integratedTerminal",
//       "program": "leetcode/383/383.go"
//     }
//   ]
// }
