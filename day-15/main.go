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
	warehouse := Warehouse{
		Moves:   []Move{},
		MapSize: [2]int{},
		Robot:   [2]int{},
		Boxes:   map[[2]int]any{},
		Walls:   map[[2]int]any{},
	}

	sc := bufio.NewScanner(strings.NewReader(input))
	l := 0
	for sc.Scan() {
		t := sc.Text()

		if t == "" {
			warehouse.MapSize = [2]int{l, warehouse.MapSize[1]}
			continue
		}

		if t[0] == '#' {
			warehouse.MapSize = [2]int{l, len(t)}
		}

		for c, r := range t {
			if r == '#' {
				warehouse.Walls[[2]int{l, c}] = nil
				continue
			}

			if r == 'O' {
				warehouse.Boxes[[2]int{l, c}] = nil
				continue
			}

			if r == '@' {
				warehouse.Robot = [2]int{l, c}
				continue
			}

			if r == '.' {
				continue
			}

			warehouse.Moves = append(warehouse.Moves, Move(r))
		}

		l++
	}

	w := warehouse.Simulate()
	println(w.BoxesGpsSum())
}

type Warehouse struct {
	Moves   []Move
	MapSize [2]int
	Robot   [2]int
	Boxes   map[[2]int]any
	Walls   map[[2]int]any
}

func (w Warehouse) Simulate() Warehouse {
	for _, m := range w.Moves {
		v := [2]int{0, 0}

		switch m {
		case MoveUp:
			v[0]--
		case MoveDown:
			v[0]++
		case MoveLeft:
			v[1]--
		case MoveRight:
			v[1]++
		}

		p := [2]int{w.Robot[0] + v[0], w.Robot[1] + v[1]}
		if _, ok := w.Walls[p]; ok {
			continue
		}

		if _, ok := w.Boxes[p]; ok {
			np := [2]int{p[0] + v[0], p[1] + v[1]}
			if _, ok2 := w.Walls[np]; ok2 {
				continue
			}

			bm := [][2]int{p}
			canMove := true
			for {
				np = bm[len(bm)-1]
				np2 := [2]int{np[0] + v[0], np[1] + v[1]}
				if _, ok2 := w.Walls[np2]; ok2 {
					canMove = false
					break
				}

				if _, ok2 := w.Boxes[np2]; ok2 {
					bm = append(bm, np2)
				} else {
					break
				}
			}

			if canMove {
				for _, b := range bm {
					delete(w.Boxes, b)
				}

				for _, b := range bm {
					np := [2]int{b[0] + v[0], b[1] + v[1]}
					w.Boxes[np] = nil
				}

				w.Robot = p
				continue
			}
		} else {
			w.Robot = p
		}

	}

	return w
}

func (w Warehouse) BoxesGpsSum() int {
	sum := 0
	for b := range w.Boxes {
		sum += 100*b[0] + b[1]
	}

	return sum
}

type Move rune

const (
	MoveUp    Move = '^'
	MoveDown  Move = 'v'
	MoveLeft  Move = '<'
	MoveRight Move = '>'
)

func mustInt(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	if c, err := strconv.ParseInt(s, 10, 64); err == nil {
		return int(c)
	}

	panic(fmt.Sprintf("%v is not int", s))
}
