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
			if string(data[i][j]) == "^" {
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
	var visited = make(map[Position]bool)
	visited[Position{x: start.x, y: start.y, dir: start.dir}] = true

	for i, line := range data {
		maze[i] = []rune(line)
	}

	currentDir := directions[start.dir]

	for {
		start.x = start.x + currentDir.x
		start.y = start.y + currentDir.y

		if start.x >= len(maze[0])-1 || start.y >= len(maze)-1 {
			visited[Position{x: start.x, y: start.y, dir: "X"}] = true
			break
		}

		if string(maze[start.x-1][start.y]) == "#" && currentDir.dir == "^" {
			currentDir = directions[">"]
			start.dir = currentDir.dir
		}

		if string(maze[start.x][start.y+1]) == "#" && currentDir.dir == ">" {
			currentDir = directions["v"]
			start.dir = currentDir.dir
		}

		if string(maze[start.x+1][start.y]) == "#" && currentDir.dir == "v" {
			currentDir = directions["<"]
			start.dir = currentDir.dir
		}

		if string(maze[start.x][start.y-1]) == "#" && currentDir.dir == "<" {
			currentDir = directions["^"]
			start.dir = currentDir.dir
		}

		if string(maze[start.x][start.y]) == "." {
			visited[Position{x: start.x, y: start.y, dir: "X"}] = true
		}

	}

	fmt.Printf("part1 : %d\n", len(visited))

}
