package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

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
