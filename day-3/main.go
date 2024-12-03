package main

import (
	"bufio"
	"flag"
	"regexp"
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
	r := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|don't\\(\\)|do\\(\\)")
	muls := [][2]int64{}
	muls2 := [][2]int64{}

	sc := bufio.NewScanner(strings.NewReader(input))
	accepts := true
	for sc.Scan() {
		t := sc.Text()
		g := r.FindAllStringSubmatch(t, -1)

		for _, sg := range g {
			if strings.HasPrefix(sg[0], "don't") {
				accepts = false
				continue
			}
			if strings.HasPrefix(sg[0], "do") {
				accepts = true
				continue
			}

			x, err := strconv.ParseInt(sg[1], 10, 64)
			if err != nil {
				panic(err)
			}

			y, err := strconv.ParseInt(sg[2], 10, 64)
			if err != nil {
				panic(err)
			}

			muls = append(muls, [2]int64{x, y})
			if accepts {
				muls2 = append(muls2, [2]int64{x, y})
			}
		}
	}

	var ts int64
	for _, mul := range muls {
		ts += mul[0] * mul[1]
	}

	var ts2 int64
	for _, mul := range muls2 {
		ts2 += mul[0] * mul[1]
	}

	println(ts)
	println(ts2)
}
