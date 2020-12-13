package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

type seat struct {
	originalState rune
	state         rune
	nextState     rune
}

type seatMap [][]*seat

const (
	empty    = 'L'
	occupied = '#'
)

func adjacentOccupiedSeats(sm seatMap, X, Y int) int {
	count := 0

	for y := Y - 1; y <= Y+1; y++ {
		for x := X - 1; x <= X+1; x++ {
			// Assume all line have the same length
			if x >= 0 &&
				x < len(sm[0]) &&
				y >= 0 &&
				y < len(sm) &&
				(x != X || y != Y) && // Skip the seat we're looking around of
				sm[y][x].state == occupied {
				count++
			}
		}
	}

	return count
}

func visibleOccupiedSeats(sm seatMap, X, Y int) int {
	count := 0

	height := len(sm)
	width := len(sm[0])

	// horizontal right
	for x := X - 1; x >= 0; x-- {
		if sm[Y][x].state == occupied {
			count++
			break
		}
	}

	// horizontal left
	for x := X + 1; x < width; x++ {
		if sm[Y][x].state == occupied {
			count++
			break
		}
	}

	// vertical top
	for y := Y - 1; y >= 0; y-- {
		if sm[y][X].state == occupied {
			count++
			break
		}
	}

	// vertical bottom
	for y := Y + 1; y < height; y++ {
		if sm[y][X].state == occupied {
			count++
			break
		}
	}

	// top left diagonal
	for i := 1; X-i >= 0 && Y-i >= 0; i++ {
		if sm[Y-i][X-i].state == occupied {
			count++
			break
		}
	}

	// top right diagonal
	for i := 1; X+i < width && Y-i >= 0; i++ {
		if sm[Y-i][X+i].state == occupied {
			count++
			break
		}
	}

	// bottom left diagonal
	for i := 1; X-i >= 0 && Y+i < height; i++ {
		if sm[Y+i][X-i].state == occupied {
			count++
			break
		}
	}

	// bottom right diagonal
	for i := 1; X+i < width && Y+i < height; i++ {
		if sm[Y+i][X+i].state == occupied {
			count++
			break
		}
	}

	log.Printf("count: %d", count)

	return count
}

func applyNextState(l []*seat) {
	for _, s := range l {
		s.state = s.nextState
	}
}

func run(sm seatMap, occupiedFunc func(seatMap, int, int) int, tolerance int) int {
	var totalOccupied int

	for {
		changed := false
		totalOccupied = 0

		for y := 0; y < len(sm); y++ {
			applyNextState(sm[y])

			if y+1 < len(sm) {
				applyNextState(sm[y+1])
			}

			for x := 0; x < len(sm[0]); x++ {
				seatState := sm[y][x].state

				if seatState == empty || seatState == occupied {
					count := occupiedFunc(sm, x, y)

					if seatState == empty && count == 0 {
						sm[y][x].nextState = occupied
						totalOccupied++
						changed = true
					}

					if seatState == occupied {
						if count >= tolerance {
							sm[y][x].nextState = empty
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

	return totalOccupied
}

func resetSeatMap(sm seatMap) {
	for _, l := range sm {
		for _, s := range l {
			s.nextState = s.originalState
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	eof := false

	sm := make(seatMap, 0)
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
			sm = append(sm, line)
			line = make([]*seat, 0)
		} else {
			line = append(line, &seat{originalState: c})
		}
	}

	resetSeatMap(sm)

	log.Printf("Part 1: %d occupied seats", run(sm, adjacentOccupiedSeats, 4))

	resetSeatMap(sm)

	log.Printf("Part 2: %d occupied seats", run(sm, visibleOccupiedSeats, 5))
}
