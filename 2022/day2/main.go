package main

import (
	"fmt"
	"os"
	"strings"
)

// Win = 6
// Draw = 3
// Lose = 1

func main() {
	c := readInput("input")

	fmt.Println("== Day 2 - Part 1 ==")
	total, x := parse(c, 1)
	fmt.Printf("Line score: %v\n", x)
	fmt.Printf("Total: %d \n", total)

	fmt.Println("")
	fmt.Println("== Day 2 - Part 2 ==")
	total, x = parse(c, 2)
	fmt.Printf("Line score: %v\n", x)
	fmt.Printf("Total: %d \n", total)
}

func readInput(file string) []string {
	c, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error load input file")
		return nil
	}

	return strings.Split(string(c), "\n")
}

func parse(c []string, part int) (int, []int) {
	x := []int{}
	total := 0

	for _, l := range c {
		if len(l) == 0 {
			continue
		}
		r := strings.Replace(l, " ", "", 1)
		f := string(r[0])
		p := string(r[1])


		score := 0
		fmt.Printf("Line: " + r + " " + f + " " + p)

		// in case of part 2 we modify the playing line
		if part == 2 {
			r = play(f, p)
			fmt.Printf(" -> %s\n", r)
		} else {
			fmt.Println("")
		}
    // re-calc in case of part 2
		p = string(r[1])

    if p == "X" {
      score += 1
    } else if p == "Y" {
      score += 2
    } else if p == "Z" {
      score += 3
    }

		// IMPROVED
		score += points(r)

		x = append(x, score)
		total += score
	}

	return total, x
}

// PART 1
// A = rock
// B = paper
// C = scissor

// X = Rock = 1
// Y = Paper = 2
// Z = Scissor 3
func points(r string) int {
	table := map[string]int{
		"AY": 6,
		"AX": 3,
		"BZ": 6,
		"BY": 3,
		"CX": 6,
		"CZ": 3,
	}

	score := 0
	for k, v := range table {
		if k == r {
			score += v
		}
	}

	// FIRST SOLUTION
	// if ot[0] == "A" {
	// 	if ot[1] == "Y" {
	// 		score += 6
	// 		fmt.Println("Won!")
	// 	} else if ot[1] == "X" {
	// 		score += 3
	// 		fmt.Println("Draw!")
	// 	}
	// } else if ot[0] == "B" {
	// 	if ot[1] == "Z" {
	// 		score += 6
	// 		fmt.Println("Won!")
	// 	} else if ot[1] == "Y" {
	// 		score += 3
	// 		fmt.Println("Draw!")
	// 	}
	// } else if ot[0] == "C" {
	// 	if ot[1] == "X" {
	// 		score += 6
	// 		fmt.Println("Won!")
	// 	} else if ot[1] == "Z" {
	// 		score += 3
	// 		fmt.Println("Draw!")
	// 	}
	// }

	return score
}

// PART 2
// A = rock
// B = paper
// C = scissor

// X = Rock = 1
// Y = Paper = 2
// Z = Scissor 3

// X = lose
// Y = draw
// Z = won
func play(p string, f string) string {
	wins := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	draws := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	loses := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	if f == "Y" {
		fmt.Printf(", Draw: " + f + " = " + draws[p])
		return p + draws[p]
	} else if f == "Z" {
		fmt.Printf(", Wins: " + f + " = " + wins[p])
		return p + wins[p]
	}

	fmt.Printf(", Lose: " + f + " = " + loses[p])
	return p + loses[p]
}
