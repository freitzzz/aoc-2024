package main

import (
	"bufio"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

//go:embed input-test.txt
var input_test string

func init() {
	flag.BoolFunc("test", "uses input-test.txt", func(s string) error {
		input = input_test
		return nil
	})

	flag.Parse()
}

func main() {
	rules := map[int][]int{}
	printPages := [][]int{}

	mustInt := func(s string) int {
		if c, err := strconv.ParseInt(s, 10, 64); err == nil {
			return int(c)
		}

		panic(fmt.Sprintf("%s is not int", s))
	}

	mustInts := func(s []string) []int {
		ints := []int{}
		for _, s := range s {
			ints = append(ints, mustInt(s))
		}

		return ints
	}

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		t := sc.Text()

		// 1st section
		if strings.Contains(t, "|") {
			split := strings.Split(t, "|")
			l := mustInt(split[0])
			r := mustInt(split[1])
			if rs, ok := rules[l]; ok {
				rs = append(rs, r)
				rules[l] = rs
				continue
			}

			rules[l] = []int{r}
			continue
		}

		// 2nd section

		if strings.Contains(t, ",") {
			printPages = append(printPages, mustInts(strings.Split(t, ",")))
		}
	}

	orderedUpdates := [][]int{}
	unorderedUpdates := [][]int{}
	for _, pages := range printPages {
		slices.Reverse(pages)
		bag := []int{}
		accepts := true

		for i, page := range pages {
			rs, ok := rules[page]
			if !ok && i == 0 {
				bag = append(bag, page)
				continue
			}

			for _, p := range bag {
				if !slices.Contains(rs, p) {
					accepts = false
					break
				}
			}

			if !accepts {
				break
			}

			bag = append(bag, page)
		}

		if accepts {
			slices.Reverse(bag)
			orderedUpdates = append(orderedUpdates, bag)
		} else {
			slices.Reverse(pages)
			unorderedUpdates = append(unorderedUpdates, pages)
		}
	}

	reorderedUpdates := [][]int{}
	for _, pages := range unorderedUpdates {
		slices.Reverse(pages)
		bag := []int{}

		for _, page := range pages {
			accepts := true

			rs, ok := rules[page]
			if !ok {
				bag = slices.Insert(bag, 0, page)
				continue
			}

			for i, p := range bag {
				if !slices.Contains(rs, p) {
					accepts = false
					bag = slices.Insert(bag, i+1, p)
					bag[i] = page

					break
				}
			}

			if accepts {
				bag = append(bag, page)
			}
		}

		slices.Reverse(bag)
		reorderedUpdates = append(reorderedUpdates, bag)
	}

	var orderedSum int64
	for _, updates := range orderedUpdates {
		orderedSum += int64(updates[len(updates)/2])
	}

	var reorderedSum int64
	for _, updates := range reorderedUpdates {
		reorderedSum += int64(updates[len(updates)/2])
	}

	println(orderedSum)
	println(reorderedSum)
}
