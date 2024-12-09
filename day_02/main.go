package main

import (
	"fmt"
	fetch "getdata"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2024/day/2/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

var sample string = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func convertToInt(nums []string) ([]int, error) {
	var converted []int
	for _, v := range nums {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("error converting '%s' to int: %v\n", v, err)
			return nil, err
		}
		converted = append(converted, num)
	}
	return converted, nil
}

func isSafe(nums []int) bool {
	increase := nums[0] > nums[1]
	for i := range len(nums) - 1 {
		diff := nums[i] - nums[i+1]
		if increase {
			if diff != 1 && diff != 2 && diff != 3 {
				return false
			}
		} else {
			if diff != -1 && diff != -2 && diff != -3 {
				return false
			}
		}
	}
	return true
}

func tolerateCheck(nums []int) bool {
	if isSafe(nums) {
		return true
	}

	for i := range len(nums) {
		var tmp []int
		left := nums[:i]
		right := nums[i+1:]
		tmp = append(tmp, left...)
		tmp = append(tmp, right...)

		if isSafe(tmp) {
			return true
		}
	}
	return false
}

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")

	var count, count2 int

	for _, v := range data {
		parts := strings.Fields(v)
		nums, err := convertToInt(parts)
		if err != nil {
			fmt.Println(err)
		}

		//part1
		if isSafe(nums) {
			count++
		}

		//part2
		if tolerateCheck(nums) {
			count2++
		}

	}
	fmt.Println("part1", count)
	fmt.Println("part2", count2)
}
