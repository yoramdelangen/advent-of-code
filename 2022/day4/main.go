package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("== DAY 4 ==")

	file, _ := os.ReadFile("input")
	lines := strings.Split(string(file), "\n")

  anyOverlap := 0
	fullOverlap := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
    fmt.Println("")

		groups := strings.Split(line, ",")
		ranges := getRanges(groups)

		fullOverlap += countFullOverlap(ranges)

    fmt.Println(hasAnyOverlap(ranges))
    if hasAnyOverlap(ranges) {
      anyOverlap += 1
    }
	}

	fmt.Println("== PART 1 ==")
	fmt.Println("Count full sections overlap:")
	fmt.Println(fullOverlap)
	fmt.Println("")

	fmt.Println("== PART 2 ==")
	fmt.Println("Count any sections overlap:")
	fmt.Println(anyOverlap)
	fmt.Println("")
}

func between(a, b, min, max int) bool {
	return a >= min && b <= max
}

func getRanges(groups []string) [][]int {
	ranges := [][]int{}

	for _, group := range groups {
		sp := strings.Split(group, "-")

		v1, _ := strconv.Atoi(sp[0])
		v2, _ := strconv.Atoi(sp[1])

		ranges = append(ranges, []int{v1, v2})
	}

	return ranges
}

func countFullOverlap(ranges [][]int) int {
	overlapping := 0
	// check for each range if it overlaps
	fmt.Println(ranges)
	if between(ranges[0][0], ranges[0][1], ranges[1][0], ranges[1][1]) {
		fmt.Println("FIRST section")
		overlapping += 1
	} else if between(ranges[1][0], ranges[1][1], ranges[0][0], ranges[0][1]) {
		overlapping += 1
		fmt.Println("SECOND section")
	}

	return overlapping
}

func inbetween(x, min, max int) bool {
	return x >= min && x <= max
}

func hasAnyOverlap(ranges [][]int) bool {
	return inbetween(ranges[0][0], ranges[1][0], ranges[1][1]) ||
		inbetween(ranges[0][1], ranges[1][0], ranges[1][1]) ||
		inbetween(ranges[1][0], ranges[0][0], ranges[0][1]) ||
		inbetween(ranges[1][1], ranges[0][0], ranges[0][1])
}
