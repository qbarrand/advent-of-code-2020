package main

import (
	"container/ring"
	"errors"
	"fmt"
	"io"
	"log"
)

const (
	N        = 25
	NotFound = -1
)

func main() {
	i := 0

	r := ring.New(N)
	m := make(map[int]int)

	part1 := NotFound

	for {
		var v int

		if _, err := fmt.Scanf("%d", &v); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("Line %d: %v", i, err)
			}

			break
		}

		if i >= N {
			canSum := false

			for p := r.Next(); p != r; p = p.Next() {
				if m[v-p.Value.(int)] > 0 {
					canSum = true
					break
				}
			}

			if !canSum {
				part1 = v
				break
			}

			m[r.Value.(int)]--
		}

		r.Value = v
		m[v]++

		r = r.Prev()
		i++
	}

	log.Printf("Part 1: cannot sum to %d", part1)
}
