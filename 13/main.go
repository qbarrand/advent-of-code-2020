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

const ignored = -1

type bus struct {
	id    int
	index int
}

func main() {
	var earliest int

	if _, err := fmt.Scanf("%d", &earliest); err != nil {
		log.Fatalf("Could not read the earliest timestamp: %v", err)
	}

	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Could not read bus IDs: %v", err)
	}

	buses := make([]*bus, 0)

	var id int

	for idx, idStr := range strings.Split(s, ",") {
		if idStr == "x" {
			continue
		}

		if id, err = strconv.Atoi(idStr); err != nil {
			log.Fatalf("Could not parse bus ID %q: %v", idStr, err)
		}

		buses = append(buses, &bus{id: id, index: idx})
	}

	soonestID := math.MaxInt64
	soonestWaitTime := math.MaxInt64

	for _, b := range buses {
		last := (earliest / b.id) * b.id
		next := last + b.id
		wait := next - earliest

		if wait < soonestWaitTime {
			soonestWaitTime = wait
			soonestID = b.id
		}
	}

	log.Printf("Part 1: %d", soonestID*soonestWaitTime)

	//n := 1
	//
	//for _, b := range buses {
	//	n = n*b.id + b.index
	//}
	//
	////n /= buses[0].id
	//
	//log.Printf("n=%d", n)

	maxID := 0

	for _, b := range buses {
		if b.id > maxID {
			maxID = b.id
		}
	}

	log.Printf("Max ID: %d", maxID)

	t := 0
	previous := 1

	for bi := 0; bi < len(buses)-1; bi++ {
		b := buses[bi]
		bnext := buses[bi+1]

		for ; ; t += b.index {
			if t%previous == 0 && t%bnext.id == b.index {
				previous = t
			}
		}
	}

	log.Printf("Previous: %d", previous)

	for t = 0; ; t += buses[0].id {
		//aligned := true

		if t >= 1068781 {
			log.Printf("t=%d", t)

			for _, b := range buses {
				log.Printf("bus %d/%d, modulo %d", b.index, b.id, (t+b.index)%b.id)

				//if t%b.id != b.index {
				//	if t == 1068781 {
				//		log.Fatalf("t=%d: bus %d/%d not aligned (modulo=%d)", t, b.index, b.id, t%b.id)
				//		return
				//	}
				//
				//	aligned = false
				//	break
				//}
			}
			return
		}

		//if aligned {
		//	break
		//}
	}

	log.Printf("Part 2: %d", t)
}
