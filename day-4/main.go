package main

import (
	"bufio"
	"flag"
	"slices"
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
	chars := [][]rune{}

	word := []rune("XMAS")
	revWord := []rune("SAMX")

	word2 := []rune("MAS")
	revWord2 := []rune("SAM")

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		t := sc.Text()
		chars = append(chars, []rune(t))
	}

	rs := len(chars)

	var count int64
	for i := 0; i < rs; i++ {
		cs := len(chars[i])
		for j := 0; j < cs; j++ {
			if r := chars[i][j]; r != 'X' && r != 'S' {
				continue
			}

			// horizontal >
			if j+3 < cs {
				if w := []rune{chars[i][j], chars[i][j+1], chars[i][j+2], chars[i][j+3]}; slices.Equal(w, word) || slices.Equal(w, revWord) {
					count++
				}
			}

			// vertical down
			if i+3 < rs {
				if w := []rune{chars[i][j], chars[i+1][j], chars[i+2][j], chars[i+3][j]}; slices.Equal(w, word) || slices.Equal(w, revWord) {
					count++
				}
			}

			// diagonal >
			if i+3 < rs && j+3 < cs {
				if w := []rune{chars[i][j], chars[i+1][j+1], chars[i+2][j+2], chars[i+3][j+3]}; slices.Equal(w, word) || slices.Equal(w, revWord) {
					count++
				}
			}

			// diagonal <
			if i-3 >= 0 && j+3 < cs {
				if w := []rune{chars[i][j], chars[i-1][j+1], chars[i-2][j+2], chars[i-3][j+3]}; slices.Equal(w, word) || slices.Equal(w, revWord) {
					count++
				}
			}
		}
	}

	var count2 int64
	for i := 0; i < rs; i++ {
		cs := len(chars[i])
		for j := 0; j < cs; j++ {
			if r := chars[i][j]; r != 'M' && r != 'S' {
				continue
			}

			// diagonal >
			if i+2 < rs && j+2 < cs {
				if w := []rune{chars[i][j], chars[i+1][j+1], chars[i+2][j+2]}; slices.Equal(w, word2) || slices.Equal(w, revWord2) {
					// diagonal <
					if w := []rune{chars[i][j+2], chars[i+1][j+1], chars[i+2][j]}; slices.Equal(w, word2) || slices.Equal(w, revWord2) {
						count2++
					}
				}
			}
		}
	}

	println(count)
	println(count2)
}
