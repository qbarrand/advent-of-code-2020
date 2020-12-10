package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

const (
	N        = 25
	NotFound = -1
)

func canSumTo(elems [N]int, sum int) bool {
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if elems[i]+elems[j] == sum {
				return true
			}
		}
	}

	return false
}

func main() {
	i := 0

	numbers := [N]int{}

	part1 := NotFound

	for {
		var v int

		if _, err := fmt.Scanf("%d", &v); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("Line %d: %v", i, err)
			}

			break
		}

		if i >= N && !canSumTo(numbers, v) {
			part1 = v
			break
		}

		numbers[i%N] = v

		i++
	}

	log.Printf("Part 1: cannot sum to %d", part1)
}
