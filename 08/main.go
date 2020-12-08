package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var errLoop = errors.New("loop")

type instruction struct {
	operation string
	argument  int
}

func main() {
	instructions := make([]*instruction, 0)
	i := 0

	for {
		in := instruction{}

		if _, err := fmt.Fscanf(os.Stdin, "%s %d", &in.operation, &in.argument); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("Line %d: %v", i, err)
			}

			break
		}

		instructions = append(instructions, &in)

		i++
	}

	log.Printf("Part 1: acc: %d", findLoop(instructions))

	acc2, err := fixProgram(instructions, 0, make([]bool, len(instructions)), true)
	if err != nil {
		log.Fatal("Could not fix the program :(")
	}

	log.Printf("Part 2: acc: %d", acc2)
}

// part 1
func findLoop(instructions []*instruction) int {
	executed := make([]bool, len(instructions))

	acc := 0
	i := 0

	for i < len(instructions) {
		in := instructions[i]

		if executed[i] {
			break
		}

		executed[i] = true

		switch in.operation {
		case "acc":
			acc += in.argument
		case "jmp":
			i += in.argument
			continue
		case "nop":
			// Do nothing
		default:
			log.Fatalf("Instruction %d: unhandled operation %q", i, in.operation)
		}

		i++
	}

	return acc
}

// part 2
func fixProgram(instructions []*instruction, i int, executed []bool, changeAllowed bool) (int, error) {
	if executed[i] {
		return 0, errLoop
	}

	if i == len(instructions)-1 {
		return 0, nil
	}

	executed[i] = true

	switch in := instructions[i]; in.operation {
	case "acc":
		n, err := fixProgram(instructions, i+1, cloneSlice(executed), changeAllowed)
		return in.argument + n, err
	case "jmp":
		if changeAllowed {
			if n, err := fixProgram(instructions, i+1, cloneSlice(executed), false); err == nil {
				return n, err
			}

			return fixProgram(instructions, i+in.argument, cloneSlice(executed), true)
		}

		return fixProgram(instructions, i+in.argument, cloneSlice(executed), false)
	case "nop":
		if changeAllowed {
			if n, err := fixProgram(instructions, i+in.argument, cloneSlice(executed), false); err == nil {
				return n, err
			}

			return fixProgram(instructions, i+1, cloneSlice(executed), true)
		}

		return fixProgram(instructions, i+1, cloneSlice(executed), false)
	default:
		return 0, fmt.Errorf("%d: operation %q not understood", i, in.operation)
	}
}

func cloneSlice(s []bool) []bool {
	c := make([]bool, len(s))

	copy(c, s)

	return c
}
