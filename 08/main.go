package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

var errLoop = errors.New("loop")

type instruction struct {
	argument  int
	executed  bool
	operation string
}

func (i *instruction) reset() {
	i.executed = false
}

func main() {
	instructions := make([]*instruction, 0)
	i := 0

	for {
		in := instruction{}

		if _, err := fmt.Scanf("%s %d", &in.operation, &in.argument); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("Line %d: %v", i, err)
			}

			break
		}

		instructions = append(instructions, &in)

		i++
	}

	log.Printf("Part 1: acc: %d", findLoop(instructions))

	for _, in := range instructions {
		in.reset()
	}

	acc2, err := fixProgram(instructions, 0, true)
	if err != nil {
		log.Fatal("Could not fix the program :(")
	}

	log.Printf("Part 2: acc: %d", acc2)
}

// part 1
func findLoop(instructions []*instruction) int {
	acc := 0
	i := 0

	for i < len(instructions) {
		in := instructions[i]

		if in.executed {
			break
		}

		in.executed = true

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
func fixProgram(instructions []*instruction, i int, changeAllowed bool) (int, error) {
	in := instructions[i]

	if in.executed {
		return 0, errLoop
	}

	if i == len(instructions)-1 {
		return 0, nil
	}

	in.executed = true
	defer in.reset()

	switch in.operation {
	case "acc":
		n, err := fixProgram(instructions, i+1, changeAllowed)
		return in.argument + n, err
	case "jmp":
		if changeAllowed {
			if n, err := fixProgram(instructions, i+1, false); err == nil {
				return n, err
			}
		}

		return fixProgram(instructions, i+in.argument, changeAllowed)
	case "nop":
		if changeAllowed {
			if n, err := fixProgram(instructions, i+in.argument, false); err == nil {
				return n, err
			}
		}

		return fixProgram(instructions, i+1, changeAllowed)
	default:
		return 0, fmt.Errorf("%d: operation %q not understood", i, in.operation)
	}
}
