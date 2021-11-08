package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"sort"
)

func main() {
	adapters := make([]int, 0)

	for i := 1; ; i++ {
		var j int

		if _, err := fmt.Scanf("%d", &j); err != nil {
			if !errors.Is(io.EOF, err) {
				log.Fatalf("Line %d: %v", i, err)
			}

			break
		}

		adapters = append(adapters, j)
	}

	sort.Ints(adapters)

	diff1 := 1 // charging outlet
	diff3 := 1 // device

	if len(adapters) >= 2 {
		for i := 1; i < len(adapters); i++ {
			switch adapters[i] - adapters[i-1] {
			case 1:
				diff1++
			case 3:
				diff3++
			}
		}
	}

	log.Printf("Part 1: %d", diff1*diff3)
}
