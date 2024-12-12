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

var cacheFn = map[[2]int]int{}

func init() {
	flag.BoolFunc("test", "uses input-test.txt", func(s string) error {
		input = input_test
		return nil
	})

	flag.Parse()
}

func main() {
	farm := [][]rune{}
	plots := map[rune][]Region{}

	sc := bufio.NewScanner(strings.NewReader(input))
	l := 0
	for sc.Scan() {
		t := sc.Text()
		farm = append(farm, []rune{})

		for _, r := range t {
			farm[l] = append(farm[l], r)

			if _, ok := plots[r]; !ok {
				plots[r] = []Region{}
			}
		}

		l++
	}

	for l := 0; l < len(farm); l++ {
		for c := 0; c < len(farm[l]); c++ {
			if plant := farm[l][c]; plant != '*' {
				var region Region

				farm, region = discoverRegion(farm, l, c)
				plots[plant] = append(plots[plant], region)

			}
		}
	}

	println(part1(plots))
}

func part1(plots map[rune][]Region) int {
	sum := 0
	for _, rs := range plots {
		for _, r := range rs {
			sum += r.Price()
		}
	}

	return sum
}

func discoverRegion(farm [][]rune, l, c int) ([][]rune, Region) {
	region := Region{Positions: [][2]int{}}
	positions := map[[2]int]any{{l, c}: nil}
	plant := farm[l][c]

	maxl := len(farm)
	maxc := len(farm[0])

	checkpoints := [][2]int{{l, c}}
	for {
		if len(checkpoints) == 0 {
			break
		}

		l = checkpoints[0][0]
		c = checkpoints[0][1]

		// mark as seen
		farm[l][c] = '*'

		// left
		if c-1 >= 0 && farm[l][c-1] == plant {
			p := [2]int{l, c - 1}
			if _, ok := positions[p]; !ok {
				positions[p] = nil
				checkpoints = append(checkpoints, p)
			}
		}

		// right
		if c+1 < maxc && farm[l][c+1] == plant {
			p := [2]int{l, c + 1}
			if _, ok := positions[p]; !ok {
				positions[p] = nil
				checkpoints = append(checkpoints, p)
			}
		}

		// up
		if l-1 >= 0 && farm[l-1][c] == plant {
			p := [2]int{l - 1, c}
			if _, ok := positions[p]; !ok {
				positions[p] = nil
				checkpoints = append(checkpoints, p)
			}

		}

		// down
		if l+1 < maxl && farm[l+1][c] == plant {
			p := [2]int{l + 1, c}
			if _, ok := positions[p]; !ok {
				positions[p] = nil
				checkpoints = append(checkpoints, p)
			}
		}

		checkpoints = checkpoints[1:]
	}

	for p := range positions {
		region.Positions = append(region.Positions, p)
	}

	return farm, region
}

type Region struct {
	Positions [][2]int
}

func (r Region) Area() int {
	return len(r.Positions)
}

func (r Region) Perimeter() int {
	sum := 0
	pcache := map[[2]int]any{}
	for _, p := range r.Positions {
		pcache[p] = nil
	}

	for _, p := range r.Positions {
		l := p[0]
		c := p[1]

		// left
		if _, ok := pcache[[2]int{l, c - 1}]; !ok {
			sum++
		}

		// right
		if _, ok := pcache[[2]int{l, c + 1}]; !ok {
			sum++
		}

		// up
		if _, ok := pcache[[2]int{l - 1, c}]; !ok {
			sum++
		}

		// down
		if _, ok := pcache[[2]int{l + 1, c}]; !ok {
			sum++
		}
	}

	return sum
}

func (r Region) Price() int {
	return r.Area() * r.Perimeter()
}
