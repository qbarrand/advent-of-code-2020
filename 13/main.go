package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var earliest int

	if _, err := fmt.Scanf("%d", &earliest); err != nil {
		log.Fatalf("Could not read the earliest timestamp: %v", err)
	}

	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Could not read bus IDs: %v", err)
	}

	ids := make([]int, 0)

	var id int

	for _, idStr := range strings.Split(s, ",") {
		if idStr != "x" {
			if id, err = strconv.Atoi(idStr); err != nil {
				log.Fatalf("Could not parse bus ID %q: %v", idStr, err)
			}

			ids = append(ids, id)
		}
	}

	soonestID := math.MaxInt64
	soonestWaitTime := math.MaxInt64

	for _, i := range ids {
		last := (earliest / i) * i
		next := last + i
		wait := next - earliest

		if wait < soonestWaitTime {
			soonestWaitTime = wait
			soonestID = i
		}
	}

	log.Printf("Part 1: %d", soonestID*soonestWaitTime)
}
