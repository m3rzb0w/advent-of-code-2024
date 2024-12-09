package main

import (
	"fmt"
	fetch "getdata"
	"sort"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2024/day/1/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

var sample = `
3   4
4   3
2   5
1   3
3   9
3   3
`

var left []int
var right []int

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")
	for _, v := range data {
		parts := strings.Fields(v)

		numLeft, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Printf("error converting '%s' to int: %v\n", parts[0], err)
		}

		numRight, err := strconv.Atoi(parts[1])

		if err != nil {
			fmt.Printf("error converting '%s' to int: %v\n", parts[1], err)
		}

		left = append(left, numLeft)
		right = append(right, numRight)
	}

	sort.Ints(left)
	sort.Ints(right)

	//part1
	var count int

	for i := range left {
		if left[i] < right[i] {
			count += right[i] - left[i]
		} else {
			count += left[i] - right[i]
		}
	}

	fmt.Println("part1", count)

	//part2
	counters := make(map[int]int)

	for i := range left {
		for j := range right {
			if left[i] == right[j] {
				counters[left[i]]++
			}
		}
	}

	var count2 int

	for k, v := range counters {
		if v == 1 {
			count2 += k
		} else {
			count2 += k * v
		}
	}

	fmt.Println("part2", count2)
}
