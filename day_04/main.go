package main

import (
	"fmt"
	fetch "getdata"
	"strings"
)

var url string = "https://adventofcode.com/2024/day/4/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

var sample string = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

type Position struct {
	x, y int
}

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")
	fmt.Println(sample)

	//part1
	var poslist = []Position{{1, 0}, {0, -1}, {-1, 0}, {0, 1}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	var count int

	for i := range data {
		for j := range data[i] {
			if string(data[i][j]) != "X" {
				continue
			}
			for _, v := range poslist {
				maxX := j + 3*v.x
				maxY := i + 3*v.y
				if !(maxY < len(data) && maxX < len(data[0]) && 0 <= maxX && 0 <= maxY) {
					continue
				}
				if string(data[i+v.y][j+v.x]) == "M" && string(data[i+2*v.y][j+2*v.x]) == "A" && string(data[i+3*v.y][j+3*v.x]) == "S" {
					count++
				}
			}

		}

	}
	fmt.Printf("part1 : %d\n", count)

	//part2
	poslist = []Position{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	var count2 int
	for i := 1; i < len(data)-1; i++ {
		for j := 1; j < len(data[0])-1; j++ {
			if string(data[i][j]) != "A" {
				continue
			}
			var tmp string
			for _, v := range poslist {
				tmp += string(data[i+v.y][j+v.x])
			}
			switch tmp {
			case "MMSS", "MSSM", "SSMM", "SMMS":
				count2++
			default:
				fmt.Println("no xmas for you")
			}
		}
	}
	fmt.Printf("part2 : %d\n", count2)
}
