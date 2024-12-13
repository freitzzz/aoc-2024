package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"strconv"
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
	machines := []Machine{}

	sc := bufio.NewScanner(strings.NewReader(input))
	m := Machine{}
	for sc.Scan() {
		t := sc.Text()

		if t == "" {
			continue
		}

		t = strings.ReplaceAll(t, "+", "=")
		t = strings.ReplaceAll(t, "X=", "")
		t = strings.ReplaceAll(t, "Y=", "")

		split := strings.Split(t, ":")
		split2 := strings.Split(split[1], ",")

		if strings.HasPrefix(split[0], "Button A") {
			m.A = [2]int{mustInt(split2[0]), mustInt(split2[1])}
			continue
		}

		if strings.HasPrefix(split[0], "Button B") {
			m.B = [2]int{mustInt(split2[0]), mustInt(split2[1])}
			continue
		}

		if strings.HasPrefix(split[0], "Prize") {
			m.Prize = [2]int{mustInt(split2[0]), mustInt(split2[1])}

			machines = append(machines, m)
			m = Machine{}
			continue
		}
	}

	machines2 := []Machine{}
	for _, m := range machines {
		m.Prize = [2]int{m.Prize[0] + 10000000000000, m.Prize[1] + 10000000000000}
		machines2 = append(machines2, m)
	}

	println(part1(machines))
	println(part2(machines2))
}

func part1(
	machines []Machine,
) int {
	sum := 0
	for _, m := range machines {

		if tk := tokens(m.A, m.B, m.Prize); tk != -1 {
			sum += tk
		}
	}

	return sum
}

func part2(
	machines []Machine,
) int {
	sum := 0
	for _, m := range machines {
		if tk := tokens2(m.A, m.B, m.Prize); tk != -1 {
			sum += tk
		}
	}

	return sum
}

func tokens(
	a, b, p [2]int,
) int {
	ca, cb := solve(a[0], a[1], b[0], b[1], p[0], p[1])

	if ca == cb && ca == 0 {
		return -1
	}

	return 3*ca + cb
}

func tokens2(
	a, b, p [2]int,
) int {
	ca, cb := solve2(a[0], a[1], b[0], b[1], p[0], p[1])

	if ca == cb && ca == 0 {
		return -1
	}

	return 3*ca + cb
}

func solve(
	ax, ay, bx, by, px, py int,
) (int, int) {
	a := 100

	for {
		if a == -1 {
			return 0, 0
		}

		b := (px - ax*a) / bx
		bmod := (px - ax*a) % bx
		if bmod != 0 {
			a--
			continue
		}

		b2 := (py - ay*a) / by
		bmod2 := (py - ay*a) % by
		if bmod2 != 0 {
			a--
			continue
		}

		if b == b2 {
			return a, b
		}

		a--
	}
}

// obrigado zé pedro, rute, hyperneutrino e o mano que inventou o teorema do mod chines
func solve2(
	ax, ay, bx, by, px, py int,
) (int, int) {
	ca := float64(px*by-py*bx) / float64(ax*by-ay*bx)
	cb := (float64(px) - float64(ax)*ca) / float64(bx)

	if math.Mod(ca, 1) == 0 && math.Mod(cb, 1) == 0 {
		return int(ca), int(cb)
	}

	return 0, 0
}

type Machine struct {
	A     [2]int
	B     [2]int
	Prize [2]int
}

func mustInt(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	if c, err := strconv.ParseInt(s, 10, 64); err == nil {
		return int(c)
	}

	panic(fmt.Sprintf("%v is not int", s))
}
