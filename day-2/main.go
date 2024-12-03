package main

import (
	"bufio"
	"flag"
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
	rr := [][]int64{}

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		s := strings.Fields(sc.Text())
		l := []int64{}
		for _, r := range s {
			n, err := strconv.ParseInt(r, 10, 64)
			if err != nil {
				panic(err)
			}

			l = append(l, n)
		}

		rr = append(rr, l)
	}

	safe := func(x, y int64, inc bool) bool {
		var diff int64
		if inc {
			diff = y - x
		} else {
			diff = x - y
		}

		return diff >= 1 && diff <= 3
	}

	var ts int64
	c := len(rr)
	for i := 0; i < c; i++ {
		l := rr[i]
		inc := false
		s := true
		for j := 0; j < len(l); j++ {
			if j == 0 {
				continue
			}

			if j == 1 && l[1]-l[0] >= 0 {
				inc = true
			}

			if !safe(l[j-1], l[j], inc) {
				s = false
				break
			}
		}

		if s {
			ts++
		}
	}

	var tsd int64
	for i := 0; i < c; i++ {
		l := rr[i]
		inc := false
		outl := 0
		for j := 0; j < len(l); j++ {
			if j == 0 {
				continue
			}

			if j == 1 && l[1]-l[0] >= 0 {
				inc = true
			}

			if safe(l[j-1], l[j], inc) {
				continue
			}

			if j+1 == len(l) {
				continue
			}

			outl++
			if !safe(l[j-1], l[j+1], inc) {
				outl++
				break
			}

			if j-2 > 0 && !safe(l[j-2], l[j], inc) {
				outl++
				break
			}
		}

		if outl <= 1 {
			tsd++
		}
	}

	println(ts)
	println(tsd)
}
