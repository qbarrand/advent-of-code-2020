package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	reMask = regexp.MustCompile(`^mask = ([X01]+)$`)
	reMem  = regexp.MustCompile(`^mem\[(\d+)] = (\d+)$`)
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var currentMask string

	m := make(map[int]int)

	for i := 0; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			log.Fatalf("Line %d: %v", i, err)
		}

		line := sc.Text()

		if strings.HasPrefix(line, "mask") {
			matches := reMask.FindStringSubmatch(line)
			currentMask = matches[1]
		} else if strings.HasPrefix(line, "mem") {
			matches := reMem.FindStringSubmatch(line)

			idx, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("Could not convert memory index %q to integer: %v", matches[1], err)
			}

			valInt, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Fatalf("Could not convert memory index %q to integer: %v", matches[2], err)
			}

			val := strconv.FormatInt(int64(valInt), 2)

			strings.TrimLeft()

			n := 0

			for idx, c := range val {
				if c ==  currentMask[idx] == '1'

				n += int(math.Pow(2, float64(idx)+1))
			}

			_ = idx
			_ = val

			m[idx] = n
		} else {
			log.Fatalf("Cannot parse line %d: %q", i, line)
		}
	}
}
