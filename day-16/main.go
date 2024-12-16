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

	paths := allPaths(maze.Map, [][][2]int{}, [][2]int{maze.StartPosition}, [][2]int{maze.StartPosition}, maze.StartPosition[0], maze.StartPosition[1])
	fmt.Printf("paths: %v\n", paths)
}

type Maze struct {
	Map           [][]rune
	StartPosition [2]int
	EndPosition   [2]int
}

func (m Maze) FindBestPath() [][2]int {
	maxx := len(m.Map)
	maxy := len(m.Map[0])
	mazeMap := m.Map

	p := m.StartPosition
	x := p[0]
	y := p[1]

	// 13, 1
	// > 13, 2
	// > 12, 1

	paths := [][][2]int{}
	allCheckpoints := [][][2]int{}

	allCheckpoints = append(allCheckpoints, [][2]int{p})
	for {
		if len(allCheckpoints) == 0 {
			break
		}

		checkpoints := allCheckpoints[len(allCheckpoints)-1]
		x = checkpoints[len(checkpoints)-1][0]
		y = checkpoints[len(checkpoints)-1][1]

		if mazeMap[x][y] == 'E' {
			paths = append(paths, checkpoints)
			allCheckpoints = allCheckpoints[:len(allCheckpoints)-1]
			continue
		}

		// fmt.Printf("checkpoints: %v\n", checkpoints)

		added := false
		// right
		if y+1 < maxy && mazeMap[x][y+1] != '#' {
			if !contains(checkpoints, [2]int{x, y + 1}) {
				allCheckpoints = append(allCheckpoints, append(checkpoints, [2]int{x, y + 1}))
				added = true
			}
		}

		// left
		if y-1 >= 0 && mazeMap[x][y-1] != '#' {
			if !contains(checkpoints, [2]int{x, y - 1}) {
				allCheckpoints = append(allCheckpoints, append(checkpoints, [2]int{x, y - 1}))
				added = true
			}
		}

		// up
		if x-1 >= 0 && mazeMap[x-1][y] != '#' {
			if !contains(checkpoints, [2]int{x - 1, y}) {
				allCheckpoints = append(allCheckpoints, append(checkpoints, [2]int{x - 1, y}))
				added = true
			}
		}

		// down
		if x+1 < maxx && mazeMap[x+1][y] != '#' {
			if !contains(checkpoints, [2]int{x + 1, y}) {
				allCheckpoints = append(allCheckpoints, append(checkpoints, [2]int{x + 1, y}))
				added = true
			}
		}

		if !added {
			allCheckpoints = allCheckpoints[:len(allCheckpoints)-1]
			fmt.Printf("len(allCheckpoints): %v\n", len(allCheckpoints))
			continue
		}
	}

	bpl := 99999999999999999
	var bp [][2]int
	for _, p := range paths {
		if len(p) < bpl {
			bp = p
		}
	}

	return bp
}

func allPaths(maze [][]rune, all [][][2]int, paths [][2]int, visited [][2]int, x, y int) [][][2]int {
	if maze[x][y] == 'E' {
		println("FINAL")
		all = append(all, paths)
		return all
	}

	maxx := len(maze)
	maxy := len(maze[0])

	// right
	if y+1 < maxy && maze[x][y+1] != '#' {
		if !contains(visited, [2]int{x, y + 1}) {
			println("right")
			paths = append(paths, [2]int{x, y + 1})
			visited = append(visited, [2]int{x, y + 1})
			all = allPaths(maze, all, paths, visited, x, y+1)
		}
	}

	// left
	if y-1 >= 0 && maze[x][y-1] != '#' {
		if !contains(visited, [2]int{x, y - 1}) {
			println("left")
			paths = append(paths, [2]int{x, y - 1})
			visited = append(visited, [2]int{x, y - 1})
			all = allPaths(maze, all, paths, visited, x, y-1)
		}
	}

	// up
	if x-1 >= 0 && maze[x-1][y] != '#' {
		if !contains(visited, [2]int{x - 1, y}) {
			println("up")
			paths = append(paths, [2]int{x - 1, y})
			visited = append(visited, [2]int{x - 1, y})
			all = allPaths(maze, all, paths, visited, x-1, y)
		}
	}

	// down
	if x+1 < maxx && maze[x+1][y] != '#' {
		if !contains(visited, [2]int{x + 1, y}) {
			println("down")
			paths = append(paths, [2]int{x + 1, y})
			visited = append(visited, [2]int{x + 1, y})
			all = allPaths(maze, all, paths, visited, x+1, y)
		}
	}

	
	fmt.Printf("paths: %v\n", paths)
	lastP := paths[len(paths)-2]
	fmt.Printf("lastP: %v\n", lastP)
	// return allPaths(
	// 	maze,
	// 	all,
	// 	paths[0:len(paths)-1],
	// 	visited,
	// 	lastP[0],
	// 	lastP[1],
	// )
	return all
}

func contains(checkpoints [][2]int, p [2]int) bool {
	for i := len(checkpoints) - 1; i >= 0; i-- {
		if checkpoints[i][0] == p[0] && checkpoints[i][1] == p[1] {
			return true
		}
	}

	return false
}

func mustInt(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	if c, err := strconv.ParseInt(s, 10, 64); err == nil {
		return int(c)
	}

	panic(fmt.Sprintf("%v is not int", s))
}
