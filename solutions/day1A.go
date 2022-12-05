package solutions

import (
	"fmt"
	"strconv"
	"github.com/bradisbon/AoC-2022/reader"
)

func Day1A() {
	scanner := reader.Reader("inputs/day1A.txt")

	currCal := 0
	maxCal := 0

	for scanner.Scan() {
		calS := scanner.Text()
		if calS != "" {
			cal, err := strconv.Atoi(calS)
			if err != nil {
				panic(err)
			}
			currCal += cal

		} else {
			if currCal > maxCal {
				maxCal = currCal
			}
			currCal = 0
		}
	}

	fmt.Printf("Max calories: %v\n", maxCal)
}
