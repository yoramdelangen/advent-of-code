package main

import (
	"fmt"
	"os"
	"strings"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	f, _ := os.ReadFile("input")
	lines := strings.Split(string(f), "\n")

	fmt.Println("== PART 1 ==")

	letters := compartmentation(lines)

	prioSum, pp := prioritized(letters)

	fmt.Println("Letters: ")
	fmt.Println(letters)

	fmt.Println("SUM: ")
	fmt.Println(prioSum)

	fmt.Println("PRIORITIZED: ")
	fmt.Println(pp)

	fmt.Println("")
	fmt.Println("== PART 2 ==")

	letters = grouping(lines)
	prioSum, pp = prioritized(letters)

	fmt.Println("Letters: ")
	fmt.Println(letters)

	fmt.Println("SUM: ")
	fmt.Println(prioSum)

	fmt.Println("PRIORITIZED: ")
	fmt.Println(pp)
}

func InSlice(char rune, s string) bool {
	return strings.ContainsRune(s, char)
}

func compartmentation(lines []string) []string {
	letters := []string{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		// fmt.Println("")

		idx := len(line) / 2

		comp1 := line[0:idx]
		comp2 := line[idx:]

		// fmt.Println(len(line))
		// fmt.Println(len(comp1))
		// fmt.Println(len(comp2))
		// fmt.Println(line)

		found := false
		for _, l := range comp1 {
			if InSlice(l, comp2) {
				found = true
				letters = append(letters, string(l))

				// fmt.Println("Found: "+ string(l))

				break
			}
		}

		if found == false {
			fmt.Println("ERROR! Not letter found")
			break
		}

		// fmt.Println("comp1: "+ string(comp1))
		// fmt.Println("comp2: "+ string(comp2))
	}

	return letters
}

func prioritized(letters []string) (int, map[string]int) {
	prioritized := map[string]int{}
	prioSum := 0

	// lets find priority
	for _, l := range letters {
		p := strings.Index(ALPHABET, l) + 1

		fmt.Println("Letter " + string(l) + " has prio ")
		fmt.Println(p)

		prioSum += p
		prioritized[string(l)] = p
	}

	return prioSum, prioritized
}

func grouping(lines []string) []string {
	groups := []string{}

	for i := 0; i < len(lines); i = i + 3 {
		line := lines[i]
		if len(line) == 0 {
			continue
		}

		for _, character := range line {
			if strings.ContainsRune(lines[i+1], character) && strings.ContainsRune(lines[i+2], character) {
				groups = append(groups, string(character))
				fmt.Println("Found: " + string(character) + " in second(" + lines[i+1] + ") and third(" + lines[i+2] + ") line")
				break
			}

		}
		fmt.Println(groups)

	}
  return groups
}
