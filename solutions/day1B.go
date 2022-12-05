package solutions

import (
	"container/heap"
	"fmt"
	"strconv"

	"github.com/bradisbon/AoC-2022/intHeap"
	"github.com/bradisbon/AoC-2022/reader"
)

func Day1B() {
	scanner := reader.Reader("inputs/day1A.txt")

	currCal := 0
	
	top3 := intHeap.IntHeap{}
	heap.Init(&top3)

	for scanner.Scan() {
		calS := scanner.Text()
		if calS != "" {
			cal, err := strconv.Atoi(calS)
			if err != nil {
				panic(err)
			}
			currCal += cal
		} else {
			heap.Push(&top3, currCal)
			if top3.Len() > 3 {
				heap.Pop(&top3)
			}			
			currCal = 0
		}
	}
	heap.Push(&top3, currCal)
	if top3.Len() > 3 {
		heap.Pop(&top3)
	}			

	t3Sum := 0
	for len(top3) > 0 {
		t3Sum += top3.Pop().(int)
	}

	fmt.Printf("Total top 3: %v\n", t3Sum)

}
