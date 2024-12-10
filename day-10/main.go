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
	topoMap := [][]int{}
	trailheads := [][2]int{}

	sc := bufio.NewScanner(strings.NewReader(input))
	i := 0
	for sc.Scan() {
		t := sc.Text()
		topoMap = append(topoMap, []int{})
		for j := range t {
			if t[j] != '.' {
				topoMap[i] = append(topoMap[i], mustInt(t[j]))
			} else {
				topoMap[i] = append(topoMap[i], 20)
			}

			if t[j] == '0' {
				trailheads = append(trailheads, [2]int{i, j})
			}
		}

		i++
	}

	println(part1(topoMap, trailheads))
}

func part1(topoMap [][]int, trailheads [][2]int) int {
	score := map[int]int{}

	maxx := len(topoMap)
	maxy := len(topoMap[0])
	for i, th := range trailheads {
		x := th[0]
		y := th[1]

		highestCheckpoint := map[[2]int]any{}
		count := 0
		checkpoints := [][2]int{th}
		for {
			if len(checkpoints) == 0 {
				break
			}

			x = checkpoints[0][0]
			y = checkpoints[0][1]

			if topoMap[x][y] == 9 {
				count++
				checkpoints = checkpoints[1:]
				highestCheckpoint[[2]int{x, y}] = 0
				continue
			}

			checkpoints = checkpoints[1:]

			// right
			if y+1 < maxy && topoMap[x][y+1]-1 == topoMap[x][y] {
				checkpoints = slices.Insert(checkpoints, 0, [2]int{x, y + 1})
			}

			// left
			if y-1 >= 0 && topoMap[x][y-1]-1 == topoMap[x][y] {
				checkpoints = slices.Insert(checkpoints, 0, [2]int{x, y - 1})
			}

			// up
			if x-1 >= 0 && topoMap[x-1][y]-1 == topoMap[x][y] {
				checkpoints = slices.Insert(checkpoints, 0, [2]int{x - 1, y})
			}

			// down
			if x+1 < maxx && topoMap[x+1][y]-1 == topoMap[x][y] {
				checkpoints = slices.Insert(checkpoints, 0, [2]int{x + 1, y})
			}

		}

		score[i] = len(highestCheckpoint)
	}

	sum := 0
	for _, sc := range score {
		sum += sc
	}

	return sum
}

func mustInt(b byte) int {
	if c, err := strconv.ParseInt(string(b), 10, 64); err == nil {
		return int(c)
	}

	panic(fmt.Sprintf("%c is not int", b))
}
