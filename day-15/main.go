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

	// println(warehouse.Simulate().BoxesGpsSum())
	// println(warehouse.Resize().Simulate2().BoxesGpsSum2())
	fmt.Printf("Initial Robot: %v\n", warehouse.Resize().Robot)
	w := warehouse.Resize().Simulate2()
	fmt.Printf("Robot: %v\n", w.Robot)
	boxes := w.Boxes
	for b := range boxes {
		fmt.Printf("b: %v\n", b)
	}
	// walls := w.Walls
	// for w := range walls {
	// 	fmt.Printf("%v\n", w)
	// }
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

func (w Warehouse) Simulate2() Warehouse {
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

		p2 := p
		p2[1]--
		/// qq

		if _, ok := w.Boxes[p]; !ok {
			if _, ok2 := w.Boxes[p2]; ok2 {
				p3 := p2
				p2 = p
				p = p3
			} else {
				w.Robot = p
				continue
			}
		}

		if _, ok := w.Boxes[p]; ok {
			np := [2]int{p[0] + v[0], p[1] + v[1]}
			np2 := [2]int{p2[0] + v[0], p2[1] + v[1]}
			if _, ok2 := w.Walls[np]; ok2 {
				continue
			}

			if _, ok2 := w.Walls[np2]; ok2 {
				continue
			}

			bm := [][2]int{p}
			check := [][2]int{p, p2}
			canMove := true
			alreadyPassed := map[[2]int]any{}
			for {
				if len(check) == 0 {
					break
				}

				np = check[0]

				np3 := [2]int{np[0] + v[0], np[1] + v[1]}
				np4 := [2]int{np3[0], np3[1] + 1}
				np5 := [2]int{np3[0], np3[1] - 1}
				if _, ok2 := w.Walls[np3]; ok2 {
					canMove = false
					break
				}

				if _, ok2 := w.Boxes[np3]; ok2 {
					bm = append(bm, np3)

					if _, ok3 := alreadyPassed[np3]; !ok3 {
						check = append(check, np3)
					}

					if _, ok3 := alreadyPassed[np4]; !ok3 {
						check = append(check, np4)
					}
				} else if _, ok2 := w.Boxes[np5]; ok2 {
					bm = append(bm, np5)

					if _, ok3 := alreadyPassed[np3]; !ok3 {
						check = append(check, np3)
					}

					if _, ok3 := alreadyPassed[np5]; !ok3 {
						check = append(check, np5)
					}
				}

				alreadyPassed[check[0]] = nil
				check = check[1:]
			}

			if canMove {
				for _, b := range bm {
					delete(w.Boxes, b)
				}

				for _, b := range bm {
					np := [2]int{b[0] + v[0], b[1] + v[1]}
					w.Boxes[np] = nil
				}

				w.Robot = [2]int{w.Robot[0] + v[0], w.Robot[1] + v[1]}
				continue
			}
		} else {
			w.Robot = p
		}

	}

	return w
}

func (w Warehouse) Resize() Warehouse {
	w.MapSize[1] *= 2

	nwalls := map[[2]int]any{}
	for w := range w.Walls {
		nwalls[[2]int{w[0], w[1] * 2}] = nil
		nwalls[[2]int{w[0], w[1]*2 + 1}] = nil
	}

	nboxes := map[[2]int]any{}
	for b := range w.Boxes {
		nboxes[[2]int{b[0], b[1] * 2}] = nil
		// nboxes[[2]int{b[0], b[1]*2 + 1}] = nil
	}

	w.Walls = nwalls
	w.Boxes = nboxes
	w.Robot = [2]int{w.Robot[0], w.Robot[1] * 2}

	return w
}

func (w Warehouse) BoxesGpsSum() int {
	sum := 0
	for b := range w.Boxes {
		sum += 100*b[0] + b[1]
	}

	return sum
}

func (w Warehouse) BoxesGpsSum2() int {
	sum := 0
	for b := range w.Boxes {
		if _, ok := w.Boxes[[2]int{b[0], b[+1]}]; ok {
			sum += 100*b[0] + b[1]
		}
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
