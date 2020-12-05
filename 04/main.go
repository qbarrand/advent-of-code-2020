package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mapFromStringSlice(s []string) map[string]bool {
	m := make(map[string]bool, len(s))

	for _, e := range s {
		m[e] = true
	}

	return m
}

// cid temporarily not required
var (
	requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validEyeColors = mapFromStringSlice([]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"})
	reHcl          = regexp.MustCompile(`^#[a-f0-9]{6}$`)
	reHgt          = regexp.MustCompile(`^(\d+)(cm|in)$`)
	rePid          = regexp.MustCompile(`^\d{9}$`)
)

type passport map[string]string

func (p passport) Valid0() bool {
	for _, f := range requiredFields {
		if p[f] == "" {
			return false
		}
	}

	return true
}

func (p passport) Valid1() bool {
	// byr
	if byr, err := strconv.Atoi(p["byr"]); err != nil {
		return false
	} else if byr < 1920 || byr > 2002 {
		return false
	}

	// iyr
	if iyr, err := strconv.Atoi(p["iyr"]); err != nil {
		return false
	} else if iyr < 2010 || iyr > 2020 {
		return false
	}

	// eyr
	if eyr, err := strconv.Atoi(p["eyr"]); err != nil {
		return false
	} else if eyr < 2020 || eyr > 2030 {
		return false
	}

	// hgt
	m := reHgt.FindStringSubmatch(p["hgt"])

	if len(m) < 3 {
		return false
	}

	// m[0] is the full match
	number, err := strconv.Atoi(m[1])
	if err != nil {
		return false
	}

	switch m[2] {
	case "cm":
		if number < 150 || number > 193 {
			return false
		}
	case "in":
		if number < 59 || number > 76 {
			return false
		}
	default:
		return false
	}

	// hcl
	if !reHcl.MatchString(p["hcl"]) {
		return false
	}

	// ecl
	if !validEyeColors[p["ecl"]] {
		return false
	}

	// pid
	if !rePid.MatchString(p["pid"]) {
		return false
	}

	return true
}

func main() {
	valid0 := 0
	valid1 := 0

	i := 0

	var (
		err  error
		line string
	)

	eof := false
	r := bufio.NewReader(os.Stdin)

	p := passport{}

	for !eof {
		if line, err = r.ReadString('\n'); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("Line %d: %v", i, err)
			}

			eof = true
		}

		line = strings.TrimSuffix(line, "\n")

		// Try to read the line
		if line != "" {
			for _, e := range strings.Split(line, " ") {
				attr := strings.Split(e, ":")
				if len(attr) != 2 {
					log.Fatalf("Bad attribute %q line %d: %q", e, i, line)
				}

				p[attr[0]] = attr[1]
			}
		}

		// Commit if necessary
		if line == "" || eof {
			if p.Valid0() {
				valid0++
			}

			if p.Valid1() {
				valid1++
			}

			p = passport{}

			continue
		}

		i++
	}

	log.Printf("Part 1: %d valid passports", valid0)
	log.Printf("Part 2: %d valid passports", valid1)
}
