package main

import (
	"fmt"
	"os"
	"strings"
)

const MARKER_LENGTH_PART_1 = 4
const MARKER_LENGTH_PART_2 = 14

func main() {
	r, _ := os.ReadFile("input")
	lines := strings.Split(string(r), "\n")

  fmt.Println("== DAY 6 ==")
  fmt.Println("== Part 1 ==")

	for _, line := range lines {
    if len(line) == 0 {
      break
    }

		fmt.Println("Line: " + line)
    fmt.Printf("PART 1: %d\n", getPosition(line, MARKER_LENGTH_PART_1))
    fmt.Printf("PART 2: %d\n", getPosition(line, MARKER_LENGTH_PART_2))
		fmt.Println("")
  }
}

func getPosition(line string, ml int) int {
  letters := strings.Split(line, "")

	// walk through the line
  for i := ml+1; i < len(letters); i += 1 {
    b := letters[i-(ml+1):i-1]
    c := uniqueCount(b)

    // fmt.Printf("%+v count: %d", b, c)

    if c == ml {
      return i - 1
    }
	}

  return -1
}

func uniqueCount(letters []string) int {
  m := ""
  for _, l := range letters {
    if !strings.Contains(m, l) {
      m += l
    }
  }

  return len(m)
}
