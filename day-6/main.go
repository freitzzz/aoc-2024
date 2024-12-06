// many horror attempts have been commited when trying to solve part2

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
	gameMap := [][]rune{}
	obstacle := '#'
	guard := '^'

	var initialGuardPosition Position
	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		t := sc.Text()
		layer := []rune(t)
		gameMap = append(gameMap, layer)

		for i, t := range layer {
			if t == guard {
				initialGuardPosition = Position{L: len(gameMap) - 1, C: i}
				break
			}
		}
	}

	path := map[Position]int{}
	direction := Direction(0)
	currentPosition := initialGuardPosition
	path[initialGuardPosition] = 1

	maxL := len(gameMap)
	maxC := len(gameMap[0])

	for {
		var nextPosition Position
		if direction.Up() {
			nextPosition = Position{L: currentPosition.L - 1, C: currentPosition.C}
		} else if direction.Right() {
			nextPosition = Position{L: currentPosition.L, C: currentPosition.C + 1}
		} else if direction.Down() {
			nextPosition = Position{L: currentPosition.L + 1, C: currentPosition.C}
		} else if direction.Left() {
			nextPosition = Position{L: currentPosition.L, C: currentPosition.C - 1}
		}

		// finished
		if nextPosition.L < 0 || nextPosition.C < 0 || nextPosition.L >= maxL || nextPosition.C >= maxC {
			break
		}

		// turn if obstacle
		if gameMap[nextPosition.L][nextPosition.C] == obstacle {
			direction = direction.Turn()
			continue
		}

		currentPosition = nextPosition
		if _, ok := path[nextPosition]; ok {
			path[nextPosition]++
		} else {
			path[nextPosition] = 1
		}
	}

	count := len(path)

	count2 := 0
	for p := range path {
		l := p.L
		c := p.C

		r := gameMap[l][c]
		if r != obstacle && r != guard {
			m := [][]rune{}

			for i := 0; i < len(gameMap); i++ {
				m = append(m, make([]rune, len(gameMap[0])))
				for j := 0; j < len(gameMap[0]); j++ {
					m[i][j] = gameMap[i][j]
				}
			}

			m[l][c] = obstacle
			if !isSolvable(m, initialGuardPosition) {
				count2++
			}
		}

	}

	println(count)
	println(count2)
}

// same thing as part 1 but some modifications to check if the guard is stuck on a loop
func isSolvable(gameMap [][]rune, initialGuardPosition Position) bool {
	path := map[Position]int{}
	direction := Direction(0)
	currentPosition := initialGuardPosition
	path[initialGuardPosition] = 1

	maxL := len(gameMap)
	maxC := len(gameMap[0])

	stuck := 0

	for {
		var nextPosition Position
		if direction.Up() {
			nextPosition = Position{L: currentPosition.L - 1, C: currentPosition.C}
		} else if direction.Right() {
			nextPosition = Position{L: currentPosition.L, C: currentPosition.C + 1}
		} else if direction.Down() {
			nextPosition = Position{L: currentPosition.L + 1, C: currentPosition.C}
		} else if direction.Left() {
			nextPosition = Position{L: currentPosition.L, C: currentPosition.C - 1}
		}

		// finished
		if nextPosition.L < 0 || nextPosition.C < 0 || nextPosition.L >= maxL || nextPosition.C >= maxC {
			break
		}

		if gameMap[nextPosition.L][nextPosition.C] == '#' {
			if stuck >= 2 {
				return false
			}
			stuck++
			direction = direction.Turn()
			continue
		}

		currentPosition = nextPosition
		if _, ok := path[nextPosition]; ok {
			path[nextPosition]++
			if path[nextPosition] > 4 {
				return false
			}
		} else {
			path[nextPosition] = 1
		}

		stuck = 0
	}

	return true
}

type Direction int
type Position struct {
	L, C int
}

func (d Direction) Turn() Direction {
	if d.Up() {
		return 1
	}

	if d.Right() {
		return 2
	}

	if d.Down() {
		return 3
	}

	return 0
}

func (d Direction) Up() bool {
	return d == 0
}

func (d Direction) Right() bool {
	return d == 1
}

func (d Direction) Down() bool {
	return d == 2
}

func (d Direction) Left() bool {
	return d == 3
}
