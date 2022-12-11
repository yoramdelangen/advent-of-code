package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	c, _ := os.ReadFile("input")
	lines := strings.Split(string(c), "\n")

	crates, moves := parseFile(lines)

	fmt.Println("== Day5 ==")
	fmt.Println("== Part 1 ==")
	organize(crates, moves, 1)

	crates, moves = parseFile(lines)
	fmt.Println("== Part 2 ==")
	organize(crates, moves, 2)
}

type Move struct {
	Num  int
	From int
	To   int
}

type Crates = [][]string

func parseFile(lines []string) (Crates, []Move) {
	header := Crates{}
	moves := []Move{}

	hasHeader := false
	for _, line := range lines {

		if len(line) == 0 && hasHeader == true { // stop when EOF
			break
		} else if len(line) == 0 {
			hasHeader = true
			continue
		}

		// first lines are the header
		if hasHeader == false {
			// fmt.Println("")
			// fmt.Println(line)
			// skip if only numbers
			if !strings.Contains(line, "[") {
				continue
			}

			idx := 0
			for i := 0; i < len(line); i += 4 {
				container := strings.ReplaceAll(line[i:i+3], " ", "")
				container = strings.ReplaceAll(container, "[", "")
				container = strings.ReplaceAll(container, "]", "")

				if len(header) <= idx {
					header = append(header, []string{})
				}

				if len(container) == 0 {
					idx += 1
					continue
				}
				// fmt.Println(container)
				// fmt.Println(idx)

				header[idx] = append(header[idx], container)
				idx += 1
			}

			continue
		}

		re := regexp.MustCompile(`\d+`)
		f := re.FindAllString(line, -1)

		moves = append(moves, Move{
			Num:  strToInt(f[0]),
			From: strToInt(f[1]) - 1,
			To:   strToInt(f[2]) - 1,
		})
	}

	// reverse all containers
	for i := range header {
		sort.SliceStable(header[i], func(i, j int) bool {
			return i > j
		})
	}

	return header, moves
}

func strToInt(i string) int {
	o, _ := strconv.Atoi(i)
	return o
}

func organize(crates Crates, moves []Move, part int) {
	// start moving the crates around
	for _, m := range moves {
		f := crates[m.From]

		cm := f[len(f)-m.Num:]

		// reverse the string because of the natural movement of elements,
		// make sure the first item on the stack ends on top
		if part == 1 {
			sort.SliceStable(cm, func(i, j int) bool {
				return i > j
			})
		}

		// Move from => to
		crates[m.To] = append(crates[m.To], cm...)
		// Remove from items
		crates[m.From] = crates[m.From][:len(f)-m.Num]
	}

	final := []string{}
	for _, c := range crates {
		final = append(final, c[len(c)-1])
	}

	fmt.Println(strings.Join(final, ""))
}
