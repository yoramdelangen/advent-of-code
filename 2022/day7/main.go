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

			fmt.Println("Path: " + path)
		} else if strings.HasPrefix(line, "dir") {
			//
		} else {
      // when its a file size with filename
			r := regexp.MustCompile(`\d+`)
			fs, _ := strconv.Atoi(r.FindString(line))

			sizes[path] += fs

			fmt.Printf("Filesize: %T\n", fs)
			fmt.Println(fs)
		}
	}

  // only grab the filesizes directories lower then 100k
  sumSizes := 0
  for _, size := range sizes{
    if size > 100000 {
      continue
    }

    sumSizes += size
  }

	fmt.Printf("Sizes: %+v\n", sizes)
	fmt.Printf("Sum sizes at most 100k: %+v\n", sumSizes)
}
