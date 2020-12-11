package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	const (
		notFound = -1
		target   = 2020
	)

	var (
		currentInt int
		m          = make(map[int]int)
		part1      = notFound
		part2      = notFound
	)

	for i := 1; part1 == notFound || part2 == notFound; i++ {
		if _, err := fmt.Scanf("%d", &currentInt); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("Line %d: %v", i, err)
			}

			break
		}

		v1 := target - currentInt

		if part1 == notFound && m[v1] > 0 {
			part1 = v1 * currentInt
		}

		if part2 == notFound {
			m[v1]--

			for k, v := range m {
				if v2 := v1 - k; v > 0 && m[v2] > 0 {
					part2 = v1 * v2 * currentInt
				}
			}

			m[v1]++
		}

		m[currentInt]++
	}

	log.Printf("Part 1: %d", part1)
	log.Printf("Part 2: %d", part2)
}
