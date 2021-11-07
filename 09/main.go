package main

import (
	"container/ring"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
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

	var v int

	for {
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

	log.Printf("Part 1: %d", part1)

	part2 := NotFound

	if _, err := os.Stdin.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("Failed to seek(0)")
	}

	sum := 0

	r = ring.New(2)

	for i := 0; i < 2; i++ {
		if _, err := fmt.Scanf("%d", &v); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("Line %d: %v", i, err)
			}

			break
		}

		r.Value = v
		r = r.Next()

		sum += v
	}

	r = r.Prev()

	for {
		if sum == part1 {
			min := r.Value.(int)

			r.Do(func(i interface{}) {
				if ii := i.(int); ii < min {
					min = ii
				}
			})

			max := r.Value.(int)

			r.Do(func(i interface{}) {
				if ii := i.(int); ii > max {
					max = ii
				}
			})

			part2 = min + max

			break
		}

		if sum < part1 {
			// let's read more data
			if _, err := fmt.Scanf("%d", &v); err != nil {
				if !errors.Is(err, io.EOF) {
					log.Fatalf("Line %d: %v", i, err)
				}

				break
			}

			n := ring.New(1)
			n.Value = v

			// Add a link to the ring and set r to that link (most recent insertion)
			r = r.Link(n).Prev()

			sum += v

			continue
		}

		if sum > part1 {
			rem := r.Unlink(1).Value.(int)
			sum -= rem
		}
	}

	log.Printf("Part 2: %d", part2)
}
