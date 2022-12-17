package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	c, _ := os.ReadFile("input")
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

	fmt.Println("")
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
	fmt.Printf("Sum sizes at most 100k: %+v\n", sumSizes)
}

// recusively adding sizes to all its path directories
func IncrementSizes(sizes map[string]int, path string, size int) {
	sizes[path] += size

	sp := strings.Split(path, "/")
	fmt.Println(sp)
    for i := 1; i < len(sp); i++ {
		p := strings.Join(sp[:i], "/")

        sizes[p] +=  size
	}
}
