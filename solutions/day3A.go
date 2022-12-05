package solutions

import (
	"fmt"

	"github.com/bradisbon/AoC-2022/reader"
)

func CharScore(ch int) int {
	if ch > 96 {
		return ch - 96
	} else {
		return ch - 64 + 26
	}
}

func lineScore(line string) int {
	first := line[0:len(line)/2]
	second := line[len(line)/2:]

	firstChars := map[int]bool{}
	for _,ch := range first {
		firstChars[CharScore(int(ch))] = true
	}

	for _, ch := range second {
		score := CharScore(int(ch))
		if ok := firstChars[score]; ok {
			return score
		}
	}

	return 0
}

func Day3A() {
	scanner := reader.Reader("inputs/day3.txt")

	totalScore := 0
	for scanner.Scan() {
		totalScore += lineScore(scanner.Text())
	}

	fmt.Printf("total score: %v\n", totalScore)
}
