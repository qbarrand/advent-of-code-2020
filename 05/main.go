package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

const (
	rows    = 128
	columns = 8
)

func findSeat(totalRows, totalColumns int, input string) (int, int) {
	low := 0
	up := totalRows - 1
	row := 0

	for _, c := range input[:7] {
		midPoint := low + (up-low)/2

		switch c {
		case 'F':
			up = midPoint
			row = low
		case 'B':
			low = midPoint + 1
			row = up
		}
	}

	column := 0

	left := 0
	right := totalColumns - 1

	for _, c := range input[7:] {
		midPoint := left + (right-left)/2

		switch c {
		case 'L':
			right = midPoint
			column = left
		case 'R':
			left = midPoint + 1
			column = right
		}
	}

	return row, column
}

func main() {
	i := 0
	maxID := 0

	var line string
	ids := make([]int, 0)

	for {
		if _, err := fmt.Fscanln(os.Stdin, &line); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("line %d: %v", i, err)
			}

			break
		}

		r, c := findSeat(rows, columns, line)

		id := r*8 + c

		ids = append(ids, id)

		if id > maxID {
			maxID = id
		}
	}

	// Part 1
	log.Printf("Max ID: %d", maxID)

	// Part 2
	sort.Ints(ids)

	for i := 0; i < len(ids)-1; i++ {
		current := ids[i]

		if ids[i+1] == current+2 {
			log.Printf("My seat ID: %d", current+1)
			return
		}
	}

	log.Fatal("ID not found :(")
}
