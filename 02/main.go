package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type policy struct {
	n0, n1 int
}

type policy0 struct {
	*policy

	char string
}

func (p *policy0) isPasswordValid(pw string) bool {
	c := strings.Count(pw, p.char)

	return c >= p.n0 && c <= p.n1
}

type policy1 struct {
	*policy

	char rune
}

func (p *policy1) isPasswordValid(pw string) bool {
	c0 := rune(pw[p.n0-1])
	c1 := rune(pw[p.n1-1])

	return (c0 == p.char && c1 != p.char) || (c0 != p.char && c1 == p.char)
}

func main() {
	re, err := regexp.Compile(`(\d+)-(\d+)\s([[:lower:]]):\s([[:lower:]]+)`)
	if err != nil {
		log.Fatalf("Could not compile the regex: %v", err)
	}

	policy0Valid := 0
	policy1Valid := 0

	basePolicy := policy{}
	p0 := policy0{policy: &basePolicy}
	p1 := policy1{policy: &basePolicy}

	s := bufio.NewScanner(os.Stdin)

	for i := 1; s.Scan(); i++ {
		if err := s.Err(); err != nil {
			log.Fatalf("Line %d: %v", i, err)
		}

		matches := re.FindStringSubmatch(s.Text())
		if len(matches) != 5 {
			log.Fatalf("Line %d: expected 4 matches, found %d: %v", i, len(matches), matches)
		}

		// matches[0] is the full string

		if basePolicy.n0, err = strconv.Atoi(matches[1]); err != nil {
			log.Fatalf("Could not parse n0 (%q): %v", matches[1], err)
		}

		if basePolicy.n1, err = strconv.Atoi(matches[2]); err != nil {
			log.Fatalf("Could not parse n1 (%q): %v", matches[2], err)
		}

		p0.char = matches[3]
		p1.char = rune(matches[3][0])

		if p0.isPasswordValid(matches[4]) {
			policy0Valid++
		}

		if p1.isPasswordValid(matches[4]) {
			policy1Valid++
		}
	}

	log.Printf("Policy 0: %d valid passwords", policy0Valid)
	log.Printf("Policy 1: %d valid passwords", policy1Valid)
}
