package main

import (
	"bufio"
	"fmt"
	"os"
)

func containsExactlyNOccurences(codepoints map[rune]int, n int) bool {
	for _, v := range codepoints {
		if v == n {
			return true
		}
	}
	return false
}

func differingLettersCountAndIdx(word1 string, word2 string) (int, int) {
	differencesCount := 0
	var firstDifferenceIdx int

	for idx, _ := range word1 {
		if word1[idx] != word2[idx] {
			differencesCount += 1
			firstDifferenceIdx = idx
		}
	}

	return differencesCount, firstDifferenceIdx
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	exactlyTwoCount := 0
	exactlyThreeCount := 0
	var boxIds []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentId := scanner.Text()
		boxIds = append(boxIds, currentId)
		codepoints := make(map[rune]int)

		for _, letter := range currentId {
			codepoints[letter] += 1
		}

		if containsExactlyNOccurences(codepoints, 2) {
			exactlyTwoCount += 1
		}

		if containsExactlyNOccurences(codepoints, 3) {
			exactlyThreeCount += 1
		}
	}

	for _, firstWord := range boxIds {
		for _, secondWord := range boxIds {
			diffCount, _ := differingLettersCountAndIdx(firstWord, secondWord)
			if diffCount == 1 {
				fmt.Println(firstWord)
				fmt.Println(secondWord)
			}
		}
	}

	fmt.Println(
		exactlyTwoCount, "IDs with exactly two letters,", exactlyThreeCount, "with exactly three")
	fmt.Println("Checksum =", exactlyTwoCount * exactlyThreeCount)
}
