package main

import (
	"bufio"
	"flag"
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
	l := []int64{}
	r := []int64{}

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		s := strings.Fields(sc.Text())
		ln, err := strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			panic(err)
		}

		rn, err := strconv.ParseInt(s[1], 10, 64)
		if err != nil {
			panic(err)
		}

		l = append(l, ln)
		r = append(r, rn)
	}

	slices.Sort(l)
	slices.Sort(r)

	var d int64
	c := len(l)
	for i := 0; i < c; i++ {
		if l[i] <= r[i] {
			d += r[i] - l[i]
			continue
		}

		d += l[i] - r[i]
	}

	var ss int64
	for i := 0; i < c; i++ {
		for j := 0; j < c; j++ {
			n := 0
			if l[i] == r[j] {
				n++
			}

			ss += l[i] * int64(n)
		}
	}

	println(d)
	println(ss)
}
