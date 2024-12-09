package main

import (
	"fmt"
	fetch "getdata"
	"regexp"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2024/day/3/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

var sample string = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
var sample2 string = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func main() {
	input = strings.TrimSpace(input)
	//part1
	pattern := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(input, -1)
	var count int
	for _, v := range matches {
		nums := strings.Split(v[4:len(v)-1], ",")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Printf("error converting '%s' to int: %v\n", nums[0], err)
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Printf("error converting '%s' to int: %v\n", nums[1], err)
		}
		count += num1 * num2
	}

	//part2
	var count2 int
	pattern = `mul\(\d+,\d+\)|don\'t\(\)|do\(\)`
	re = regexp.MustCompile(pattern)
	matches = re.FindAllString(input, -1)
	var isEnabled bool = true
	for _, v := range matches {
		if isEnabled && strings.Contains(v, "mul") {
			nums := strings.Split(v[4:len(v)-1], ",")
			num1, err := strconv.Atoi(nums[0])
			if err != nil {
				fmt.Printf("error converting '%s' to int: %v\n", nums[0], err)
			}
			num2, err := strconv.Atoi(nums[1])
			if err != nil {
				fmt.Printf("error converting '%s' to int: %v\n", nums[1], err)
			}
			count2 += num1 * num2
		}

		if v == "don't()" {
			isEnabled = false
		}

		if v == "do()" {
			isEnabled = true
		}
	}

	fmt.Println("part1", count)
	fmt.Println("part2", count2)
}
