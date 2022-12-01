package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("== Day1 ==")
  // o := day1("input1_1.test")
  o := day1("input1_1.prod")
  p1 := o[len(o)-1]

  fmt.Printf("Part 1: %v\n", p1)

  i2 := o[len(o)-3:]

  p2 := 0
  for _, x := range i2 {
    p2 += x
  }

  fmt.Printf("Part 2: %v\n", p2)
}

func day1(input string) []int {
	// read input file
	f, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Input file was not found")
		return nil
	}

	c := string(f)
	c1 := strings.Split(c, "\n\n")

	out := []int{}

	for _, cc := range c1 {
		l := strings.Split(cc, "\n")
		// fmt.Println(l)
		t := 0
		for _, i := range l {

			// convert string into int
			n, err := strconv.Atoi(strings.TrimSpace(i))
			if err != nil {
				continue
			}

			// sum upp
			t += n
		}

		out = append(out, t)
	}

	sort.Ints(out)

  return out

}
