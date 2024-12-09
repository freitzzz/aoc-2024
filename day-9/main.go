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
	diskMap := [][2]int{}

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		t := sc.Text()
		l := len(t)
		for i := 0; i < l; i += 2 {
			f := mustInt(t[i])
			if i+1 >= l {
				diskMap = append(diskMap, [2]int{f, 0})
				continue
			}

			s := mustInt(t[i+1])
			diskMap = append(diskMap, [2]int{f, s})
		}
	}

	blocks := []int{}
	for i, blk := range diskMap {
		bf := repeat(i, blk[0])
		bs := repeat(-1, blk[1])

		blocks = append(blocks, bf...)
		blocks = append(blocks, bs...)
	}

	chks := part1(blocks)
	println(chks)
}

func part1(blocks []int) int {
	rblocks := []int{}
	l := len(blocks)
	for i, bf := range blocks {
		if i >= l {
			break
		}

		if bf != -1 {
			rblocks = append(rblocks, bf)
			continue
		}

		for {
			l--
			bf = blocks[l]
			if bf != -1 {
				rblocks = append(rblocks, bf)
				break
			}
		}
	}

	chks := 0
	for i := 0; i < len(rblocks); i++ {
		chks += i * rblocks[i]
	}

	return chks
}

func mustInt(b byte) int {
	if c, err := strconv.ParseInt(string(b), 10, 64); err == nil {
		return int(c)
	}

	panic(fmt.Sprintf("%c is not int", b))
}

func repeat(d int, c int) []int {
	r := []int{}
	for i := 0; i < c; i++ {
		r = append(r, d)
	}

	return r
}
