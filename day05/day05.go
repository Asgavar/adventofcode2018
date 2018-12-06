package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func willReact(unit1, unit2 rune) bool {
	yesAndLeftIsUpper := unicode.IsUpper(unit1) && unicode.ToLower(unit1) == unit2
	yesAndLeftIsLower := unicode.IsLower(unit1) && unicode.ToUpper(unit1) == unit2

	return yesAndLeftIsUpper || yesAndLeftIsLower
}

func nextExplosion(unitString string) string {
	for idx, _ := range unitString {
		if idx == 0 {
			continue
		}
		prev := rune(unitString[idx-1])
		curr := rune(unitString[idx])

		if willReact(prev, curr) {
			return string(prev) + string(curr)
		}
	}

	return "xD"
}

func explode(unitString string) string {
	nextPair := nextExplosion(unitString)

	if nextPair == "xD" {
		return unitString
	}

	explosionIdx := strings.Index(unitString, nextPair)
	afterExplosion := unitString[:explosionIdx] + unitString[explosionIdx+2:]

	return explode(afterExplosion)
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	units := scanner.Text()

	exploded := explode(units)

	fmt.Println("[PART I] Units left:", len(exploded))

	countsAfterRemovals := make(map[rune]int)

	for unit := 'a'; unit <= 'z'; unit++ {
		withoutThisUnit := strings.Replace(units, string(unit), "", -1)
		withoutThisUnit = strings.Replace(withoutThisUnit, string(unicode.ToUpper(unit)), "", -1)
		countsAfterRemovals[unit] = len(explode(withoutThisUnit))
	}

	minCount := 999999999999999999
	for _, v := range countsAfterRemovals {
		if v < minCount {
			minCount = v
		}
	}

	fmt.Println("[PART II]", minCount)
}
