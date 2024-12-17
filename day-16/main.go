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
	maze := Maze{
		Map: [][]rune{},
	}

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		t := sc.Text()

		for j, r := range t {
			if r == 'S' {
				maze.StartPosition = [2]int{len(maze.Map), j}
				continue
			}

			if r == 'E' {
				maze.EndPosition = [2]int{len(maze.Map), j}
				continue
			}
		}

		maze.Map = append(maze.Map, []rune(t))
	}

	// paths := allPaths(maze.Map, [][][2]int{}, [][2]int{maze.StartPosition}, [][2]int{maze.StartPosition}, maze.StartPosition[0], maze.StartPosition[1])
	// paths2 := [][][2]int{}
	// for _, p := range paths {
	// 	if p[len(p)-1] == maze.EndPosition {
	// 		paths2 = append(paths2, p)
	// 	}
	// }

	// fmt.Printf("len(paths2): %v\n", len(paths2))

	// println(part1(paths2))

	fmt.Printf("part1and2(maze): %v\n", part1(maze))

	// println("oi1")
	// tree := buildTree(maze.Map, maze.StartPosition)
	// println("oi2")
	// // fmt.Printf("%v\n", tree.Right.Right)
	// fmt.Printf("%v\n", tree.Up.Up.Right.Right)
	// // PrintTree(tree, "", "root")
	// // fmt.Printf("maze.Map: %v\n", maze.Map)
	// paths := [][][2]int{}
	// paths = allPaths2(*tree, [][2]int{}, paths)
	// paths2 := [][][2]int{}
	// for _, p := range paths {
	// 	if p[len(p)-1] == maze.EndPosition {
	// 		paths2 = append(paths2, p)
	// 	}
	// }

	// for _, p := range paths2 {
	// 	fmt.Printf("p: %v\n", p)
	// }

	// fmt.Printf("paths: %v\n", len(paths))
	// fmt.Printf("paths2: %v\n", len(paths2))
	// fmt.Printf("part1(paths): %v\n", part1(paths2))

	// // PrintTree(tree, "", "root")
}

type Maze struct {
	Map           [][]rune
	StartPosition [2]int
	EndPosition   [2]int
}

type Tree struct {
	X     int
	Y     int
	Left  *Tree
	Right *Tree
	Up    *Tree
	Down  *Tree
}

// func (m Maze) FindBestPath() [][2]int {
// 	maxx := len(m.Map)
// 	maxy := len(m.Map[0])
// 	mazeMap := m.Map

// 	p := m.StartPosition
// 	x := p[0]
// 	y := p[1]

// 	// 13, 1
// 	// > 13, 2
// 	// > 12, 1

// 	paths := [][][2]int{}
// 	allCheckpoints := [][][2]int{}

// 	allCheckpoints = append(allCheckpoints, [][2]int{p})
// 	for {
// 		if len(allCheckpoints) == 0 {
// 			break
// 		}

// 		checkpoints := allCheckpoints[len(allCheckpoints)-1]
// 		x = checkpoints[len(checkpoints)-1][0]
// 		y = checkpoints[len(checkpoints)-1][1]

// 		if mazeMap[x][y] == 'E' {
// 			paths = append(paths, checkpoints)
// 			allCheckpoints = allCheckpoints[:len(allCheckpoints)-1]
// 			continue
// 		}

// 		// fmt.Printf("checkpoints: %v\n", checkpoints)

// 		added := false
// 		// right
// 		if y+1 < maxy && mazeMap[x][y+1] != '#' {
// 			if !contains(checkpoints, [2]int{x, y + 1}) {
// 				allCheckpoints = append(allCheckpoints, append(checkpoints, [2]int{x, y + 1}))
// 				added = true
// 			}
// 		}

// 		// left
// 		if y-1 >= 0 && mazeMap[x][y-1] != '#' {
// 			if !contains(checkpoints, [2]int{x, y - 1}) {
// 				allCheckpoints = append(allCheckpoints, append(checkpoints, [2]int{x, y - 1}))
// 				added = true
// 			}
// 		}

// 		// up
// 		if x-1 >= 0 && mazeMap[x-1][y] != '#' {
// 			if !contains(checkpoints, [2]int{x - 1, y}) {
// 				allCheckpoints = append(allCheckpoints, append(checkpoints, [2]int{x - 1, y}))
// 				added = true
// 			}
// 		}

// 		// down
// 		if x+1 < maxx && mazeMap[x+1][y] != '#' {
// 			if !contains(checkpoints, [2]int{x + 1, y}) {
// 				allCheckpoints = append(allCheckpoints, append(checkpoints, [2]int{x + 1, y}))
// 				added = true
// 			}
// 		}

// 		if !added {
// 			allCheckpoints = allCheckpoints[:len(allCheckpoints)-1]
// 			fmt.Printf("len(allCheckpoints): %v\n", len(allCheckpoints))
// 			continue
// 		}
// 	}

// 	bpl := 99999999999999999
// 	var bp [][2]int
// 	for _, p := range paths {
// 		if len(p) < bpl {
// 			bp = p
// 		}
// 	}

// 	return bp
// }

// func allPaths(maze [][]rune, all [][][2]int, paths [][2]int, visited [][2]int, x, y int) [][][2]int {
// 	if maze[x][y] == 'E' {
// 		println("FINAL")
// 		all = append(all, paths)
// 		return all
// 	}

// 	maxx := len(maze)
// 	maxy := len(maze[0])

// 	// right
// 	if y+1 < maxy && maze[x][y+1] != '#' {
// 		if !contains(visited, [2]int{x, y + 1}) {
// 			println("right")
// 			paths = append(paths, [2]int{x, y + 1})
// 			visited = append(visited, [2]int{x, y + 1})
// 			all = allPaths(maze, all, paths, visited, x, y+1)
// 		}
// 	}

// 	// left
// 	if y-1 >= 0 && maze[x][y-1] != '#' {
// 		if !contains(visited, [2]int{x, y - 1}) {
// 			println("left")
// 			paths = append(paths, [2]int{x, y - 1})
// 			visited = append(visited, [2]int{x, y - 1})
// 			all = allPaths(maze, all, paths, visited, x, y-1)
// 		}
// 	}

// 	// up
// 	if x-1 >= 0 && maze[x-1][y] != '#' {
// 		if !contains(visited, [2]int{x - 1, y}) {
// 			println("up")
// 			paths = append(paths, [2]int{x - 1, y})
// 			visited = append(visited, [2]int{x - 1, y})
// 			all = allPaths(maze, all, paths, visited, x-1, y)
// 		}
// 	}

// 	// down
// 	if x+1 < maxx && maze[x+1][y] != '#' {
// 		if !contains(visited, [2]int{x + 1, y}) {
// 			println("down")
// 			paths = append(paths, [2]int{x + 1, y})
// 			visited = append(visited, [2]int{x + 1, y})
// 			all = allPaths(maze, all, paths, visited, x+1, y)
// 		}
// 	}

// 	fmt.Printf("paths: %v\n", paths)
// 	lastP := paths[len(paths)-2]
// 	fmt.Printf("lastP: %v\n", lastP)
// 	// return allPaths(
// 	// 	maze,
// 	// 	all,
// 	// 	paths[0:len(paths)-1],
// 	// 	visited,
// 	// 	lastP[0],
// 	// 	lastP[1],
// 	// )
// 	return all
// }

// func buildTree(maze [][]rune, start [2]int) *Tree {
// 	maxx := len(maze)
// 	maxy := len(maze[0])

// 	tree := Tree{
// 		X: start[0],
// 		Y: start[1],
// 	}

// 	var rec func(maze [][]rune, node *Tree)
// 	rec = func(maze [][]rune, node *Tree) {
// 		x := node.X
// 		y := node.Y

// 		if maze[x][y] == 'E' {
// 			return
// 		}

// 		// block access
// 		maze[x][y] = '#'

// 		// right
// 		if y+1 < maxy && maze[x][y+1] != '#' {
// 			node.Right = &Tree{X: x, Y: y + 1}
// 			rec(maze, node.Right)
// 		}

// 		// left
// 		if y-1 >= 0 && maze[x][y-1] != '#' {
// 			node.Left = &Tree{X: x, Y: y - 1}
// 			rec(maze, node.Left)
// 		}

// 		// up
// 		if x-1 >= 0 && maze[x-1][y] != '#' {
// 			node.Up = &Tree{X: x - 1, Y: y}
// 			rec(maze, node.Up)
// 		}

// 		// down
// 		if x+1 < maxx && maze[x+1][y] != '#' {
// 			node.Down = &Tree{X: x + 1, Y: y}
// 			rec(maze, node.Down)
// 		}

// 		maze[x][y] = '.'
// 	}

// 	rec(maze, &tree)
// 	return &tree
// }

// func allPaths2(tree Tree, path [][2]int, paths [][][2]int) [][][2]int {
// 	path = append(path, [2]int{tree.X, tree.Y})

// 	if tree.Left == nil && tree.Right == nil && tree.Up == nil && tree.Down == nil {
// 		paths = append(paths, path)
// 	}

// 	if tree.Left != nil {
// 		paths = append(allPaths2(*tree.Left, path, paths), path)
// 	}

// 	if tree.Right != nil {
// 		paths = append(allPaths2(*tree.Right, path, paths), path)
// 	}

// 	if tree.Up != nil {
// 		paths = append(allPaths2(*tree.Up, path, paths), path)
// 	}

// 	if tree.Down != nil {
// 		paths = append(allPaths2(*tree.Down, path, paths), path)
// 	}

// 	return paths
// }

// func PrintTree(node *Tree, indent string, direction string) {
// 	if node == nil {
// 		return
// 	}

// 	// Print current node
// 	fmt.Printf("%s(%d, %d) [%s]\n", indent, node.X, node.Y, direction)

// 	// Recursively print all directions
// 	newIndent := indent + "  "
// 	PrintTree(node.Left, newIndent, "Left")
// 	PrintTree(node.Right, newIndent, "Right")
// 	PrintTree(node.Up, newIndent, "Up")
// 	PrintTree(node.Down, newIndent, "Down")
// }

func part1(maze Maze) int {
	mazeMap := maze.Map
	seen := map[[4]int]any{
		{
			maze.StartPosition[0],
			maze.StartPosition[1],
			0,
			1,
		}: nil,
	}
	heap := map[[5]int]any{
		{
			0,
			maze.StartPosition[0],
			maze.StartPosition[1],
			0,
			1,
		}: nil,
	}

	popMin := func(heap map[[5]int]any) [5]int {
		var lowestK [5]int
		lowestV := 999999999999999999
		for k := range heap {
			if k[0] < lowestV {
				lowestK = k
				lowestV = k[0]
			}
		}

		delete(heap, lowestK)
		return lowestK
	}

	for {
		if len(heap) == 0 {
			break
		}

		p := popMin(heap)
		cost := p[0]
		x := p[1]
		y := p[2]
		dx := p[3]
		dy := p[4]

		if mazeMap[x][y] == 'E' {
			return cost
		}

		tries := [][5]int{
			{cost + 1, x + dx, y + dy, dx, dy},
			{cost + 1000, x, y, dy, -dx},
			{cost + 1000, x, y, -dy, dx},
		}

		for _, t := range tries {
			if mazeMap[t[1]][t[2]] != '#' {
				if _, ok := seen[[4]int{t[1], t[2], t[3], t[4]}]; !ok {
					seen[[4]int{t[1], t[2], t[3], t[4]}] = nil
					heap[t] = nil

				}
			}
		}

		println("######")

		// mazeMap[x][y] = '.'
	}

	return -1
}

func contains(checkpoints [][2]int, p [2]int) bool {
	for i := len(checkpoints) - 1; i >= 0; i-- {
		if checkpoints[i][0] == p[0] && checkpoints[i][1] == p[1] {
			return true
		}
	}

	return false
}

func (t Tree) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("(%d, %d)", t.X, t.Y))

	if t.Left != nil {
		sb.WriteString(fmt.Sprintf("\n(Left) => (%d, %d)", t.Left.X, t.Left.Y))
	}

	if t.Right != nil {
		sb.WriteString(fmt.Sprintf("\n(Right) => (%d, %d)", t.Right.X, t.Right.Y))
	}

	if t.Up != nil {
		sb.WriteString(fmt.Sprintf("\n(Up) => (%d, %d)", t.Up.X, t.Up.Y))
	}

	if t.Down != nil {
		sb.WriteString(fmt.Sprintf("\n(Down) => (%d, %d)", t.Down.X, t.Down.Y))
	}

	return sb.String()
}

func mustInt(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	if c, err := strconv.ParseInt(s, 10, 64); err == nil {
		return int(c)
	}

	panic(fmt.Sprintf("%v is not int", s))
}
