package main

import (
	"errors"
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

func countTrees(r io.ReadSeeker, s *slope) (int, error) {
	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return 0, fmt.Errorf("could not seek: %v", err)
	}

	nextX := 0
	nextY := 0

	trees := 0

	i := 0

	var line string

	for {
		if _, err := fmt.Fscanln(r, &line); err != nil {
			if !errors.Is(err, io.EOF) {
				return 0, fmt.Errorf("line %d: %v", i, err)
			}

			break
		}

		if i == nextY {
			if line[nextX%len(line)] == '#' {
				trees++
			}

			nextX += s.incX
			nextY += s.incY
		}

		i++
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

	for _, s := range slopes {
		// countTrees calls Seek(0) on the file.
		// This works only if the file is seekable - when stdin is redirected from a regular file.
		t, err := countTrees(os.Stdin, s)
		if err != nil {
			log.Fatalf("Could not count trees: %v", err)
		}

		log.Printf("%s: %d tree", s, t)

		trees *= t
	}

	log.Printf("Total: %d trees", trees)
}
