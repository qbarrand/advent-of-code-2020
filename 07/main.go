package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func makeColor(adj, color string) string {
	return adj + " " + color
}

type bagCollection map[string]map[string]int

func (bc bagCollection) containedBags(color string) int {
	total := 0

	for col, count := range bc[color] {
		total += count + bc.containedBags(col)*count
	}

	return total
}

func (bc bagCollection) containingBags(color string) int {
	total := 0

	for k, _ := range bc {
		if bc.containingBagsRecurse(color, k) {
			total++
		}
	}

	return total
}

func (bc bagCollection) containingBagsRecurse(color, recurseFrom string) bool {
	b := bc[recurseFrom]

	if b[color] != 0 {
		return true
	}

	for k, _ := range b {
		if bc.containingBagsRecurse(color, k) {
			return true
		}
	}

	return false
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	i := 0

	bags := make(bagCollection, 0)

	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatalf("Line %d: %v", i, err)
		}

		elems := strings.Split(s.Text(), " ")

		bag := make(map[string]int)

		bags[makeColor(elems[0], elems[1])] = bag

		if elems[4] != "no" {
			j := 5

			for j < len(elems) {
				n, err := strconv.Atoi(elems[j-1])
				if err != nil {
					log.Fatalf("Could not convert %s to integer: %v", elems[j-1], err)
				}

				contentColor := makeColor(elems[j], elems[j+1])

				bag[contentColor] = n

				j += 4
			}
		}

		i++
	}

	searchedColor := "shiny gold"

	log.Printf("Part 1: %d bags can contain %s bags", bags.containingBags(searchedColor), searchedColor)
	log.Printf("Part 2: %s bags can contain up to %d bags", searchedColor, bags.containedBags(searchedColor))
}
