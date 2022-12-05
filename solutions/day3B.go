package solutions

import (
	"fmt"

	"github.com/bradisbon/AoC-2022/reader"
)

func groupScore(lines []string) int {
	chars := map[int]int{}
	for _,ch := range lines[0] {
		chars[CharScore(int(ch))] = 1
	}
	for _,ch := range lines[1] {
		score := CharScore(int(ch))
		if chars[score] > 0 {
			chars[score] = 2
		}
	}
	for _,ch := range lines[2] {
		score := CharScore(int(ch))
		if chars[score] > 1 {
			return score
		}
	}

	return 0
}

func Day3B() {
	scanner := reader.Reader("inputs/day3.txt")

	totalScore := 0
	for scanner.Scan() {
		lines := []string{}
		lines = append(lines, scanner.Text())
		for i := 0; i < 2; i++ {
			scanner.Scan()
			lines = append(lines, scanner.Text())
		}
		totalScore += groupScore(lines)
	}

	fmt.Printf("total score: %v\n", totalScore)
}
