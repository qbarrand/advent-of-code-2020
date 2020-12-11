package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type slope struct {
	incX, incY int
}

func (s *slope) String() string {
	return fmt.Sprintf("right %d, down %d", s.incX, s.incY)
}

func countTrees(sc *bufio.Scanner, s *slope) (int, error) {
	nextX := 0
	nextY := 0

	trees := 0

	for i := 0; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			return 0, fmt.Errorf("line %d: %v", i, err)
		}

		line := sc.Text()

		if i == nextY {
			if line[nextX%len(line)] == '#' {
				trees++
			}

			nextX += s.incX
			nextY += s.incY
		}
	}

	return trees, nil
}

func main() {
	slopes := []*slope{
		{incX: 1, incY: 1},
		{incX: 3, incY: 1},
		{incX: 5, incY: 1},
		{incX: 7, incY: 1},
		{incX: 1, incY: 2},
	}

	trees := 1

	fd := os.Stdin

	for _, s := range slopes {
		// This works only if the file is seekable - when stdin is redirected from a regular file.
		if _, err := fd.Seek(0, io.SeekStart); err != nil {
			log.Fatalf("Could not seek: %v", err)
		}

		t, err := countTrees(bufio.NewScanner(fd), s)
		if err != nil {
			log.Fatalf("Could not count trees: %v", err)
		}

		log.Printf("%s: %d tree", s, t)

		trees *= t
	}

	log.Printf("Total: %d trees", trees)
}
