package main

import (
  "os"
  "fmt"
  "sort"
  "strings"
  "strconv"
)

func main() {
  content, err := os.ReadFile("./input1_test")
  if err != nil {
    panic("Input file does not exists")
  }

  // first split into elves
  inpt := strings.Split(string(content), "\n\n")

  // foreach elf parse their lines
  elfCalories := []int{}
  for _, e := range inpt {
    food := strings.Split(e, "\n")

    // walk through the list of calories
    totalCalories := 0
    for _, cal := range food {
      num, _ := strconv.Atoi(cal)
      // sum the calories
      totalCalories += num
    }

    // push into slice
    elfCalories = append(elfCalories, totalCalories)
  }

  sort.Ints(elfCalories)

  fmt.Println("==Part1==")
  fmt.Println(elfCalories)
  fmt.Println(elfCalories[len(elfCalories)-1])
}
