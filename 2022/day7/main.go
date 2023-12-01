package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)
const TOTAL_SPACE = 70000000
const REQUIRED_SPACE = 30000000

func main() {
	c, _ := os.ReadFile("input_test")
	lines := strings.Split(string(c), "\n")

	sizes := map[string]int{}
	path := ""

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		// detect if its a command
		if strings.HasPrefix(line, "$ ls") {
			continue
		} else if strings.HasPrefix(line, "$ cd") {
			p := strings.ReplaceAll(line, "$ cd ", "")

			if len(path) == 0 {
				path = p
			} else if p == ".." {
				s := strings.Split(path, "/")
				path = strings.Join(s[:len(s)-1], "/")
			} else {
				if path == "/" {
					path += p
				} else {
					path += "/" + p
				}
			}

			// fix path
			if !strings.HasPrefix(path, "/") {
				path = "/" + path
			}

			fmt.Println("CD: " + path)
		} else if strings.HasPrefix(line, "dir") {
			//
		} else {
			// when its a file size with filename
			r := regexp.MustCompile(`\d+`)
			fs, _ := strconv.Atoi(r.FindString(line))

			// sizes[path] += fs
			IncrementSizes(sizes, path, fs)

			fmt.Printf("Filesize: %d\n", fs)
		}
	}

    Part1(sizes)

    Part2(sizes)
}

// recusively adding sizes to all its path directories
func IncrementSizes(sizes map[string]int, path string, size int) {
	// sizes[path] += size

	sp := strings.Split(path, "/")
    for i := 1; i < len(sp); i++ {
		p := strings.Join(sp[:i+1], "/")

        sizes[p] +=  size
	}
}

func TotalUsedSpace(sizes map[string]int) int {
	sumSizes := 0
	for _, size := range sizes {
		sumSizes += size
	}

    return sumSizes
}

func Part1(sizes map[string]int) {
	fmt.Println("")
	fmt.Println("== Part 1 ==")
	fmt.Printf("Found %d folders\n", len(sizes))

	// only grab the filesizes directories lower then 100k
	sumSizes := 0
	for f, size := range sizes {
		if size >= 100000 {
			continue
		}

		fmt.Printf("folder: %s, size: %d\n", f, size)
		sumSizes += size
	}

	// fmt.Printf("Sizes: %+v\n", sizes)
	fmt.Printf("SUM SIZES AT MOST 100K: %+v\n", sumSizes)
}

func Part2(sizes map[string]int) {
	fmt.Println("")
	fmt.Println("== Part 2 ==")

    sizes = SortByValue(sizes)

    fmt.Printf("%+v\n", sizes)

    totalSize := TotalUsedSpace(sizes)

    fmt.Printf("%+v\n", totalSize)
}

func SortByValue(ob map[string]int) map[string]int {
    keys := make([]string, 0, len(ob))
    fmt.Println(ob)

    for key := range ob {
        keys = append(keys, key)
    }

    sort.SliceStable(keys, func(i, j int) bool {
        return ob[keys[i]] < ob[keys[j]]
    })

    n := make(map[string]int)
    fmt.Println(keys)

    for _, k := range keys {
        n[k] = ob[k]
    }

    fmt.Println("idohf")
    fmt.Println(n)

    return n
}
