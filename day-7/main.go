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

	unsolvable := []Equation{}
	sum := int64(0)
	for _, eq := range equations {
		tree := buildBinTree(eq.Values...)
		if solveBinTree(*tree, eq.Values[0], eq.Result) {
			sum += eq.Result
		} else {
			unsolvable = append(unsolvable, eq)
		}
	}

	sum2 := sum
	for _, eq := range unsolvable {
		tree := buildNTree(eq.Values...)
		if solveNTree(*tree, eq.Values[0], eq.Result) {
			sum2 += eq.Result
		}
	}

	println(sum)
	println(sum2)
}

type Equation struct {
	Result int64
	Values []int64
}

type BinTree struct {
	Value int64
	Left  *BinTree
	Right *BinTree
}

type NTree struct {
	Value  int64
	Left   *NTree
	Middle *NTree
	Right  *NTree
}

func buildBinTree(values ...int64) *BinTree {
	l := len(values)
	if l == 0 {
		return nil
	}

	tree := BinTree{Value: values[0]}
	tree.Left = buildBinTree(values[1:l]...)
	tree.Right = buildBinTree(values[1:l]...)

	return &tree
}

func buildNTree(values ...int64) *NTree {
	l := len(values)
	if l == 0 {
		return nil
	}

	tree := NTree{Value: values[0]}
	tree.Left = buildNTree(values[1:l]...)
	tree.Middle = buildNTree(values[1:l]...)
	tree.Right = buildNTree(values[1:l]...)

	return &tree
}

func solveBinTree(node BinTree, sum int64, result int64) bool {
	if node.Left == nil || node.Right == nil {
		return sum == result
	}

	if sum := (sum * node.Left.Value); sum <= result {
		if solveBinTree(*node.Left, sum, result) || solveBinTree(*node.Right, sum, result) {
			return true
		}
	}

	if sum := (sum + node.Right.Value); sum <= result {
		return solveBinTree(*node.Right, sum, result) || solveBinTree(*node.Left, sum, result)
	}

	return false
}

func solveNTree(node NTree, sum int64, result int64) bool {
	if node.Left == nil || node.Middle == nil || node.Right == nil {
		return sum == result
	}

	if sum := (sum + node.Right.Value); sum <= result {
		if solveNTree(*node.Right, sum, result) || solveNTree(*node.Left, sum, result) || solveNTree(*node.Middle, sum, result) {
			return true
		}
	}

	if sum := (sum * node.Left.Value); sum <= result {
		if solveNTree(*node.Left, sum, result) || solveNTree(*node.Right, sum, result) || solveNTree(*node.Middle, sum, result) {
			return true
		}
	}

	concact := mustInt(fmt.Sprintf("%d%d", sum, node.Middle.Value))
	if sum := concact; sum <= result {
		return solveNTree(*node.Left, sum, result) || solveNTree(*node.Right, sum, result) || solveNTree(*node.Middle, sum, result)
	}

	return false
}

func printBinTree(node *BinTree, level int) {
	if node == nil {
		return
	}

	// Print the right subtree first to ensure correct indentation
	printBinTree(node.Right, level+1)

	// Print the current node with appropriate indentation
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Println(node.Value)

	// Print the left subtree
	printBinTree(node.Left, level+1)
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
