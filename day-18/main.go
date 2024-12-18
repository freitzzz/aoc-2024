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

var map_size = [2]int{71, 71}

func init() {
	flag.BoolFunc("test", "uses input-test.txt", func(s string) error {
		input = input_test
		map_size = [2]int{7, 7}
		return nil
	})

	flag.Parse()
}

func main() {
	positions := [][2]int{}
	start := [2]int{0, 0}
	end := [2]int{map_size[0] - 1, map_size[1] - 1}

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		t := sc.Text()
		split := strings.Split(t, ",")
		p := [2]int{mustInt(split[0]), mustInt(split[1])}
		positions = append(positions, p)
	}

	println(part1(positions[:1024], start, end))
}

type Maze struct {
	Map           [][]rune
	StartPosition [2]int
	EndPosition   [2]int
}

func NewMaze(
	positions [][2]int,
	start [2]int,
	end [2]int,
) Maze {
	maze := Maze{StartPosition: start, EndPosition: end, Map: [][]rune{}}

	for i := 0; i < map_size[0]; i++ {
		maze.Map = append(maze.Map, []rune(strings.Repeat(".", map_size[1])))
	}

	for _, p := range positions {
		maze.Map[p[1]][p[0]] = '#'
	}

	return maze
}

func (s Maze) Print() {
	for _, l := range s.Map {
		for _, r := range l {
			print(string(r))
		}

		println()
	}
}

func part1(positions [][2]int, start [2]int, end [2]int) int {
	maze := NewMaze(positions, start, end)

	mazeMap := maze.Map
	seen := map[[4]int]any{
		{
			maze.StartPosition[0],
			maze.StartPosition[1],
			0,
			1,
		}: nil,
	}
	heap := map[[7]int]any{
		{
			0,
			maze.StartPosition[0],
			maze.StartPosition[1],
			0,
			1,
			0,
			0,
		}: nil,
	}

	popMin := func(heap map[[7]int]any) [7]int {
		var lowestK [7]int
		lowestV := 999999999999999999
		for k := range heap {
			if k[0] < lowestV {
				lowestK = k
				lowestV = k[0]
			}
		}

		delete(heap, lowestK)
		return lowestK
	}

	for {
		if len(heap) == 0 {
			break
		}

		p := popMin(heap)
		cost := p[0]
		x := p[1]
		y := p[2]
		dx := p[3]
		dy := p[4]

		if x == maze.EndPosition[0] && y == maze.EndPosition[1] {
			return cost
		}

		// p (0,0)
		// d (0,1)
		// -----
		// p (0,1) d (0,1) (right)
		// p (0,0) d (1,0) (down)
		// p (0,0) d (-1,0) (up)
		tries := [][7]int{
			{cost + 1, x + dx, y + dy, dx, dy, x, y},
			{cost + 0, x, y, dy, -dx, x, y},
			{cost + 0, x, y, -dy, dx, x, y},
		}

		for _, t := range tries {
			if t[1] < 0 || t[2] < 0 || t[1] >= map_size[0] || t[2] >= map_size[1] {
				continue
			}

			if mazeMap[t[1]][t[2]] != '#' {
				if _, ok := seen[[4]int{t[1], t[2], t[3], t[4]}]; !ok {
					seen[[4]int{t[1], t[2], t[3], t[4]}] = nil
					heap[t] = nil
				}
			}
		}
	}

	return -1
}

func mustInt(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	if c, err := strconv.ParseInt(s, 10, 64); err == nil {
		return int(c)
	}

	panic(fmt.Sprintf("%v is not int", s))
}
