package main

import (
	"errors"
	"fmt"
	fetch "getdata"
	"strings"
)

var url string = "https://adventofcode.com/2024/day/6/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

var sample string = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

type Position struct {
	x, y int
	dir  string
}

var directions = map[string]Position{
	"^": {x: -1, y: 0, dir: "^"}, // up
	">": {x: 0, y: 1, dir: ">"},  // right
	"v": {x: 1, y: 0, dir: "v"},  // down
	"<": {x: 0, y: -1, dir: "<"}, // left
}

func (p *Position) getStartPos(data []string) (Position, error) {
	for i := range data {
		for j := range data[i] {
			if data[i][j] == '^' {
				return Position{
					x:   i,
					y:   j,
					dir: "^",
				}, nil
			}
		}
	}
	return Position{}, errors.New("no start position")
}

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")

	var start Position
	start, err := start.getStartPos(data)

	if err != nil {
		fmt.Println("start position not found")
	}

	fmt.Printf("start position : %v\n", start)

	var maze = make([][]rune, len(data))

	for i, line := range data {
		maze[i] = []rune(line)
	}

	var visited = make(map[Position]bool)
	n, m := len(maze), len(maze[0])

	for {

		visited[Position{x: start.x, y: start.y, dir: "X"}] = true

		nextX, nextY := start.x+directions[start.dir].x, start.y+directions[start.dir].y

		if nextX < 0 || nextY < 0 || nextX >= n || nextY >= m {
			break
		}

		if maze[nextX][nextY] == '#' {
			switch start.dir {
			case "^":
				start.dir = ">"
			case ">":
				start.dir = "v"
			case "v":
				start.dir = "<"
			case "<":
				start.dir = "^"
			}
		} else {
			start.x, start.y = nextX, nextY
		}

	}

	fmt.Printf("part1 : %d\n", len(visited))

	// part2
	var count int
	start, _ = start.getStartPos(data)

	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] != '.' {
				continue
			}
			maze[i][j] = '#'
			if isLooping(maze, start) {
				count++
			}
			maze[i][j] = '.'
		}
	}

	fmt.Printf("part2 : %d\n", count)
}

func isLooping(maze [][]rune, start Position) bool {
	visited := make(map[Position]bool)
	n, m := len(maze), len(maze[0])

	for {
		visited[start] = true

		nextX, nextY := start.x+directions[start.dir].x, start.y+directions[start.dir].y

		// Check if out of bounds
		if nextX < 0 || nextY < 0 || nextX >= n || nextY >= m {
			return false
		}

		// If next position is a wall, change direction clockwise
		if maze[nextX][nextY] == '#' {
			switch start.dir {
			case "^":
				start.dir = ">"
			case ">":
				start.dir = "v"
			case "v":
				start.dir = "<"
			case "<":
				start.dir = "^"
			}
		} else {
			// Move to the next position
			start.x, start.y = nextX, nextY
		}

		// If we've visited this position with the same direction, it's a loop
		if visited[start] {
			return true
		}
	}
}
