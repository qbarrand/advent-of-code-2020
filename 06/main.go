package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

type set map[rune]bool

func (s set) intersect(other set) set {
	for k, _ := range s {
		if !other[k] {
			delete(s, k)
		}
	}

	return s
}

type group = map[rune]bool

func main() {
	eof := false
	i := 0
	r := bufio.NewReader(os.Stdin)

	// For part 1
	m := make(group)

	// For part 2
	var s set

	part1Yeses := 0
	part2Yeses := 0

	for !eof {
		line, err := r.ReadString('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("Line %d: %v", i, err)
			}

			eof = true
		}

		line = strings.TrimSuffix(line, "\n")

		if line != "" {
			// Process the characters
			lineSet := make(set)

			for _, c := range line {
				m[c] = true
				lineSet[c] = true
			}

			if s == nil {
				s = lineSet
			} else {
				s = s.intersect(lineSet)
			}
		}

		// EOF is at the end of a line, not the beginning of a new one
		if line == "" || eof {
			part1Yeses += len(m)
			m = make(group)

			part2Yeses += len(s)
			s = nil
		}

		i++
	}

	log.Printf("Part 1 count sum: %d", part1Yeses)
	log.Printf("Part 2 count sum: %d", part2Yeses)
}
