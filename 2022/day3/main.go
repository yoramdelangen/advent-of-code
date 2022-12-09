package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
  f, _ := os.ReadFile("input")
  lines := strings.Split(string(f), "\n")

  letters := []string{}
  alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

  for _, line := range lines {
    if len(line) == 0 {
      continue
    }
    fmt.Println("")

    idx := len(line)/2

    comp1 := line[0:idx]
    comp2 := line[idx:]

    fmt.Println(len(line))
    fmt.Println(len(comp1))
    fmt.Println(len(comp2))
    fmt.Println(line)

    found := false
    for _, l := range comp1 {
      if InSlice(l, comp2) {
        found = true
        letters = append(letters, string(l))

        fmt.Println("Found: "+ string(l))

        break
      }
    }

    if found == false {
      fmt.Println("ERROR! Not letter found")
      break
    }

    fmt.Println("comp1: "+ string(comp1))
    fmt.Println("comp2: "+ string(comp2))
  }

  fmt.Println("Letters: ")
  fmt.Println(letters)

  prioritized := map[string]int{}
  prioSum := 0

  // lets find priority
  for _, l := range letters {
    p := strings.Index(alphabet, l) + 1

    fmt.Println("Letter "+ string(l) + " has prio ")
    fmt.Println(p)

    prioSum += p
    prioritized[string(l)] = p
  }

  fmt.Println("Sum: ")
  fmt.Println(prioSum)

  fmt.Println("Prioritized: ")
  fmt.Println(prioritized)
}

func InSlice(char rune, s string) bool {
  return strings.ContainsRune(s, char)
}
