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

	// chks := part1(blocks)
	// chks2 := part2(blocks)
	// println(chks)
	// println(chks2)
	println(part2R(blocks))
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
	oob := 1
	if len(blocks)%2 == 0 {
		oob = 0
	}

	for i := 0; i < len(rblocks)-oob; i++ {
		chks += i * rblocks[i]
	}

	return chks
}

func part2R(blocks []int) int {
	rblocks := copyArray(blocks)

	acc := []int{} //block of the files that we want to move to the free space (if there is any)
	for i := len(blocks) - 1; i >= 0; i-- {
		bf := blocks[i]

		if bf != -1 && len(acc) == 0 {
			// first element
			acc = append(acc, bf)
			continue
		}

		if bf != -1 && len(acc) > 0 && bf == acc[0] {
			// other elements equal to the first in the acc
			acc = append(acc, bf)
			continue
		}

		// We found all block file of the same group
		// Now we search for the free space
		fs := 0
		fgsize := len(acc)
		for j := 0; j < i; j++ {
			if fs > 0 && fs == fgsize {
				rblocks = changeArray(rblocks, j-fs, j, acc)
				rblocks = changeArray(rblocks, i+1, i+1+fgsize, repeat(-1, fgsize))
				break
			}
			if rblocks[j] == -1 {
				fs++
				continue
			}
			if rblocks[j] != -1 && fs > 0 {
				fs = 0
			}
		}

		acc = []int{}
		if bf != -1 {
			acc = append(acc, bf)
		}
	}

	chks := 0
	for i := 0; i < len(rblocks); i++ {
		if rblocks[i] == -1 {
			continue
		}
		chks += i * rblocks[i]
	}

	return chks
}

func part2(blocks []int) int {
	cb := blocks[0]
	lastBlocks := [][]int{{}}
	l := 0
	for _, bf := range blocks {
		if bf == -1 {
			continue
		}

		if bf != cb {
			cb = bf
			l++
			lastBlocks = append(lastBlocks, []int{})
		}

		lastBlocks[l] = append(lastBlocks[l], cb)
	}

	// decreased order
	slices.Reverse(lastBlocks)

	rblocks := []int{}
	l = len(blocks)

	sc := 0
	for i := 0; i < l; i++ {
		bf := blocks[i]

		if bf == -1 {
			sc++
			continue
		}

		if bf != -1 && sc == 0 {
			rblocks = append(rblocks, bf)
			fmt.Printf("rblocks: %v\n", rblocks)
			continue
		}

		i--
		bff := []int{}
		for _, lbf := range lastBlocks {
			if len(lbf) <= sc {
				bff = lbf

				xxx := [][]int{}
				for _, lbf := range lastBlocks {
					fmt.Printf("lbf: %v\n", lbf)
					fmt.Printf("bff: %v\n", bff)
					if !slices.Equal(bff, lbf) {
						xxx = append(xxx, lbf)
					}
				}

				lastBlocks = xxx
				fmt.Printf("lastBlocks: %v\n", lastBlocks)
				println()
				break
			}
		}

		if bff == nil || len(bff) == 0 {
			for i := 0; i < sc; i++ {
				rblocks = append(rblocks, -1)
				fmt.Printf("rblocks: %v\n", rblocks)
			}

			sc = 0
		} else {
			rblocks = append(rblocks, bff...)
			fmt.Printf("rblocks: %v\n", rblocks)
			for j := i + len(bff); j < len(blocks); j++ {
				if blocks[j] == bff[0] {
					blocks[j] = -2
				}
			}

			sc = sc - len(bff)
		}

	}

	for i := range rblocks {
		if rblocks[i] == -2 {
			rblocks[i] = -1
		}
	}

	chks := 0
	for i := 0; i < len(rblocks)-2; i++ {
		if rblocks[i] == -1 {
			continue
		}

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

func changeArray(src []int, init int, end int, values []int) []int {
	r := []int{}
	j := 0

	for i := 0; i < len(src); i++ {
		if i < init || i >= end {
			r = append(r, src[i])
			continue
		}

		r = append(r, values[j])
		j++
	}
	return r
}

func copyArray(src []int) []int {
	l := len(src)
	destination := []int{}
	for i := 0; i < l; i++ {
		destination = append(destination, src[i])
	}
	return destination
}
