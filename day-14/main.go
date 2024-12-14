package main

import (
	"bufio"
	"flag"
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

//go:embed input-test.txt
var input_test string

var map_size = [2]int{101, 103}

func init() {
	flag.BoolFunc("test", "uses input-test.txt", func(s string) error {
		input = input_test
		map_size = [2]int{11, 7}
		return nil
	})

	flag.Parse()
}

func main() {
	robots := []Robot{}

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		t := sc.Text()
		t = strings.ReplaceAll(t, "p=", "")
		t = strings.ReplaceAll(t, " v=", ",")

		points := strings.Split(t, ",")
		robots = append(robots, Robot{
			InitialPosition: [2]int{mustInt(points[0]), mustInt(points[1])},
			Velocity:        [2]int{mustInt(points[2]), mustInt(points[3])},
		})
	}

	println(part1(robots))
}

func part1(robots []Robot) int {
	positions := [][2]int{}
	middle := [2]int{map_size[0] / 2, map_size[1] / 2}

	for _, r := range robots {
		p := r.FinalPosition(100)

		for {
			if p[0] >= 0 && p[0] < map_size[0] && p[1] >= 0 && p[1] < map_size[1] {
				break
			}

			if p[0] >= map_size[0] {
				p[0] = p[0] - map_size[0]
			}

			if p[0] < 0 {
				p[0] = p[0] + map_size[0]
			}

			if p[1] >= map_size[1] {
				p[1] = p[1] - map_size[1]
			}

			if p[1] < 0 {
				p[1] = p[1] + map_size[1]
			}
		}

		positions = append(positions, p)
	}

	quadrants := [4]int{0, 0, 0, 0}
	for _, p := range positions {
		// Q1
		if p[0] < middle[0] && p[1] < middle[1] {
			quadrants[0] = quadrants[0] + 1
			continue
		}

		// Q2
		if p[0] > middle[0] && p[1] < middle[1] {
			quadrants[1] = quadrants[1] + 1
			continue
		}

		// Q3
		if p[0] < middle[0] && p[1] > middle[1] {
			quadrants[2] = quadrants[2] + 1
			continue
		}

		// Q4
		if p[0] > middle[0] && p[1] > middle[1] {
			quadrants[3] = quadrants[3] + 1
			continue
		}
	}

	factor := 1
	for _, q := range quadrants {
		if q != 0 {
			factor *= q
		}
	}

	return factor
}

type Robot struct {
	InitialPosition [2]int
	Velocity        [2]int
}

func (r Robot) FinalPosition(i int) [2]int {
	// y = ax + b
	x := linear(r.Velocity[0], r.InitialPosition[0], i)
	y := linear(r.Velocity[1], r.InitialPosition[1], i)

	return [2]int{x, y}
}

// y = ax + b
func linear(a, b, x int) int {
	return a*x + b
}

func mustInt(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	if c, err := strconv.ParseInt(s, 10, 64); err == nil {
		return int(c)
	}

	panic(fmt.Sprintf("%v is not int", s))
}
