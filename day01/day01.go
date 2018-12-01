package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	answerToPartOne := 0
	answerToPartTwo := 0
	frequencyNow := 0
	alreadyOccuredFrequencies := make(map[int]bool)
	var frequencies []int

	for scanner.Scan() {
		currentFrequencyChange, _ := strconv.Atoi(scanner.Text())
		answerToPartOne += currentFrequencyChange
		frequencies = append(frequencies, currentFrequencyChange)
	}

	stillLookingForAnswerToPartTwo := true

	for stillLookingForAnswerToPartTwo {
		for _, frequencyChange := range frequencies {
			frequencyNow += frequencyChange
			_, ok := alreadyOccuredFrequencies[frequencyNow]
			if ok && stillLookingForAnswerToPartTwo {
				answerToPartTwo = frequencyNow
				stillLookingForAnswerToPartTwo = false
			} else {
				alreadyOccuredFrequencies[frequencyNow] = true
			}
		}
	}

	fmt.Println("The resulting frequency should be", answerToPartOne)
	fmt.Println("The first frequency reached twice was", answerToPartTwo)
}
