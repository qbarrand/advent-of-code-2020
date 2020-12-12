package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type seat struct {
	state     rune
	nextState rune
}

const (
	empty    = 'L'
	occupied = '#'
)

func adjacentOccupiedSeats(l [][]*seat, X, Y int) int {
	count := 0

	for y := Y - 1; y <= Y+1; y++ {
		for x := X - 1; x <= X+1; x++ {
			// Assume all line have the same length
			if x >= 0 &&
				x < len(l[0]) &&
				y >= 0 &&
				y < len(l) &&
				(x != X || y != Y) && // Skip the seat we're looking around of
				l[y][x].state == occupied {
				count++
			}
		}
	}

	return count
}

func applyNextState(l []*seat) {
	for _, s := range l {
		s.state = s.nextState
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	eof := false

	lines := make([][]*seat, 0)
	line := make([]*seat, 0)

	for i := 0; !eof; i++ {
		c, _, err := r.ReadRune()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("Line %d: %v", i, err)
			}

			eof = true
		}

		if c == '\n' || eof {
			lines = append(lines, line)
			line = make([]*seat, 0)
		} else {
			line = append(line, &seat{state: c, nextState: c})
		}
	}

	for _, l := range lines {
		for _, s := range l {
			fmt.Printf("%c", s.state)
		}

		fmt.Print("\n")
	}

	var totalOccupied int

	for {
		changed := false
		totalOccupied = 0

		for y := 0; y < len(lines); y++ {
			applyNextState(lines[y])

			if y+1 < len(lines) {
				applyNextState(lines[y+1])
			}

			for x := 0; x < len(lines[0]); x++ {
				seatState := lines[y][x].state

				if seatState == empty || seatState == occupied {
					count := adjacentOccupiedSeats(lines, x, y)

					if seatState == empty && count == 0 {
						lines[y][x].nextState = occupied
						totalOccupied++
						changed = true
					}

					if seatState == occupied {
						if count >= 4 {
							lines[y][x].nextState = empty
							changed = true
						} else {
							totalOccupied++
						}
					}
				}
			}
		}

		if !changed {
			break
		}
	}

	log.Printf("Total occupied: %d", totalOccupied)
}
