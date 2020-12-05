package main

import (
	"bufio"
	"errors"
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

	ids := make([]int, 0)
	buf := make([]byte, 10)

	r := bufio.NewReader(os.Stdin)

	for {
		// Read the 10 characters per line
		if _, err := io.ReadFull(r, buf); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("line %d: %v", i, err)
			}

			break
		}

		// Read the new line
		if _, err := r.ReadByte(); err != nil && !errors.Is(err, io.EOF) {
			log.Fatalf("Could not read the new line: %v", err)
		}

		rows, columns := findSeat(rows, columns, string(buf))

		id := rows*8 + columns

		ids = append(ids, id)

		if id > maxID {
			maxID = id
		}

		i++
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
