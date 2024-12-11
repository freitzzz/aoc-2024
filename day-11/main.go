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

func init() {
	flag.BoolFunc("test", "uses input-test.txt", func(s string) error {
		input = input_test
		return nil
	})

	flag.Parse()
}

func main() {
	stones := []int{}

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		t := sc.Text()
		stones = append(stones, mustInts(strings.Fields(t))...)
	}

	println(part1(stones, 25))
}

func part1(stones []int, count int) int {
	for i := 0; i < count; i++ {
		ns := []int{}
		for _, st := range stones {
			if st == 0 {
				ns = append(ns, 1)
				continue
			}

			if s := strconv.Itoa(st); len(s)%2 == 0 {
				split := []int{
					mustInt(s[:len(s)/2]),
					mustInt(s[len(s)/2:]),
				}

				ns = append(ns, split...)
				continue
			}

			ns = append(ns, st*2024)
		}

		stones = ns
	}

	return len(stones)
}

func mustInt(s string) int {
	if c, err := strconv.ParseInt(s, 10, 64); err == nil {
		return int(c)
	}

	panic(fmt.Sprintf("%v is not int", s))
}

func mustInts(s []string) []int {
	ints := []int{}
	for _, s := range s {
		ints = append(ints, mustInt(s))
	}

	return ints
}
