package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	rows    = 128
	columns = 8
)

func findSeat(totalRows, totalColumns int, input string) (int, int) {
	return 0, 0
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Argument required")
	}

	inputFilename := os.Args[1]

	fd, err := os.Open(inputFilename)
	if err != nil {
		log.Fatalf("Could not open %q: %v", inputFilename, err)
	}
	defer fd.Close()

	i := 0
	max := 0

	var line string

	for {
		if _, err := fmt.Fscanln(fd, &line); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("line %d: %v", i, err)
			}

			break
		}

		r, c := findSeat(rows, columns, line)

		s := r*8 + c

		if s > max {
			max = s
		}
	}

	log.Printf("Max: %d", max)
}
