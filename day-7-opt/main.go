// https://www.youtube.com/watch?v=pSqvQiqOVO0

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
	equations := []Equation{}

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		t := sc.Text()

		split := strings.Split(t, ":")
		equation := Equation{
			Result: mustInt(split[0]),
			Values: mustInts(strings.Fields(split[1])),
		}

		equations = append(equations, equation)
	}

	sum := int64(0)
	sum2 := int64(0)
	for _, eq := range equations {
		if isValid(eq.Result, eq.Values) {
			sum += eq.Result
			sum2 += eq.Result
		} else if isValid2(eq.Result, eq.Values) {
			sum2 += eq.Result
		}
	}

	println(sum)
	println(sum2)
}

func isValid(result int64, values []int64) bool {
	if len(values) == 1 {
		return values[0] == result
	}

	if values[0] > result {
		return false
	}

	v := []int64{values[0] + values[1]}
	if isValid(result, append(v, values[2:]...)) {
		return true
	}

	v = []int64{values[0] * values[1]}
	return isValid(result, append(v, values[2:]...))
}

func isValid2(result int64, values []int64) bool {
	if len(values) == 1 {
		return values[0] == result
	}

	if values[0] > result {
		return false
	}

	v := []int64{values[0] + values[1]}
	if isValid2(result, append(v, values[2:]...)) {
		return true
	}

	v = []int64{values[0] * values[1]}
	if isValid2(result, append(v, values[2:]...)) {
		return true
	}

	v = []int64{mustInt(fmt.Sprintf("%d%d", values[0], values[1]))}
	if isValid2(result, append(v, values[2:]...)) {
		return true
	}

	return false
}

type Equation struct {
	Result int64
	Values []int64
}

func mustInt(s string) int64 {
	if c, err := strconv.ParseInt(s, 10, 64); err == nil {
		return c
	}

	panic(fmt.Sprintf("%s is not int", s))
}

func mustInts(s []string) []int64 {
	ints := []int64{}
	for _, s := range s {
		ints = append(ints, mustInt(s))
	}

	return ints
}
