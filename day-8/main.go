package main

import (
	"bufio"
	"flag"
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
	antennas := map[rune][][2]int{}

	sc := bufio.NewScanner(strings.NewReader(input))
	i := 0
	var t string
	for sc.Scan() {
		t = sc.Text()
		for j, r := range t {
			if r == '.' {
				continue
			}

			if loc, ok := antennas[r]; ok {
				antennas[r] = append(loc, [2]int{i, j})
			} else {
				antennas[r] = [][2]int{{i, j}}

			}
		}

		i++
	}

	maxl := i
	maxc := len(t)

	count := part1(antennas, maxl, maxc)
	count2 := part2(antennas, maxl, maxc)

	println(count)
	println(count2)
}

func part1(antennas map[rune][][2]int, maxl, maxc int) int {
	antinodes := [][2]int{}
	for _, locs := range antennas {
		for i := 0; i < len(locs); i++ {
			loc1 := locs[i]

			for j := i + 1; j < len(locs); j++ {
				loc2 := locs[j]
				diffl := loc2[0] - loc1[0]
				diffc := loc2[1] - loc1[1]

				an1 := [2]int{loc1[0] - diffl, loc1[1] - diffc}
				an2 := [2]int{loc2[0] + diffl, loc2[1] + diffc}

				antinodes = append(antinodes, an1, an2)
			}
		}
	}

	mapAntiNodes := map[any]any{}
	for _, an := range antinodes {
		if an[0] < 0 || an[1] < 0 || an[0] >= maxl || an[1] >= maxc {
			continue
		}

		mapAntiNodes[an] = nil
	}

	count := len(mapAntiNodes)
	return count
}

func part2(antennas map[rune][][2]int, maxl, maxc int) int {
	antinodes := [][2]int{}
	for _, locs := range antennas {
		for i := 0; i < len(locs); i++ {
			loc1 := locs[i]

			for j := i + 1; j < len(locs); j++ {
				loc2 := locs[j]
				diffl := loc2[0] - loc1[0]
				diffc := loc2[1] - loc1[1]

				an1 := [2]int{loc1[0] - diffl, loc1[1] - diffc}
				an2 := [2]int{loc2[0] + diffl, loc2[1] + diffc}
				antinodes = append(antinodes, an1, an2, loc1, loc2)

				for {
					an1 = [2]int{an1[0] - diffl, an1[1] - diffc}
					if an1[0] < 0 || an1[1] < 0 {
						break
					}

					antinodes = append(antinodes, an1)
				}

				for {
					an2 = [2]int{an2[0] + diffl, an2[1] + diffc}
					if an2[0] >= maxl || an2[1] >= maxc {
						break
					}

					antinodes = append(antinodes, an2)
				}
			}
		}
	}

	mapAntiNodes := map[any]any{}
	for _, an := range antinodes {
		if an[0] < 0 || an[1] < 0 || an[0] >= maxl || an[1] >= maxc {
			continue
		}

		mapAntiNodes[an] = nil
	}

	count := len(mapAntiNodes)
	return count
}
