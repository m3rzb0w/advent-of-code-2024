package main

import (
	"fmt"
	fetch "getdata"
	"strconv"
	"strings"
)

var url string = "https://adventofcode.com/2024/day/5/input"
var fileName = fetch.Filename(url)
var fileExists = fetch.CheckFileExists(fileName)
var session = fetch.Grabsession()
var input = fetch.Getfile(fileExists, fileName, url, session)

var sample string = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

type orderingRule struct {
	pageLeft  int
	pageRight int
}

var orderingRules []orderingRule
var pagesUpdates [][]int

func extractPageNums(nums string) orderingRule {
	splitNums := strings.Split(nums, "|")
	l, _ := strconv.Atoi(splitNums[0])
	r, _ := strconv.Atoi(splitNums[1])
	return orderingRule{
		pageLeft:  l,
		pageRight: r,
	}
}

func extractPageUpdates(nums string) []int {
	var updates []int
	splitNums := strings.Split(nums, ",")
	for _, v := range splitNums {
		n, _ := strconv.Atoi(v)
		updates = append(updates, n)
	}
	return updates
}

func main() {
	input = strings.TrimSpace(input)
	data := strings.Split(input, "\n")
	for _, v := range data {
		if strings.Contains(v, "|") {
			nums := extractPageNums(v)
			orderingRules = append(orderingRules, nums)
		}

		if strings.Contains(v, ",") {
			update := extractPageUpdates(v)
			pagesUpdates = append(pagesUpdates, update)
		}
	}

	var updatesIsOK bool
	var count, count2 int
	for _, v := range pagesUpdates {
		updatesIsOK = true
		for i := range len(v) - 1 {
			for j := range orderingRules {
				if v[i] == orderingRules[j].pageRight && orderingRules[j].pageLeft == v[i+1] {
					updatesIsOK = false
				}
			}
		}
		if updatesIsOK {
			count += v[len(v)/2]
		} else {
			//part2
			sorted := sortUpdates(v)
			count2 += sorted[len(sorted)/2]

		}
	}
	fmt.Println("part1", count)
	fmt.Println("part2", count2)
}

func sortUpdates(nums []int) []int {
	for {
		var isSorted = true
		for i := range len(nums) - 1 {
			for j := range orderingRules {
				if nums[i] == orderingRules[j].pageRight && orderingRules[j].pageLeft == nums[i+1] {
					isSorted = false
					nums[i] = orderingRules[j].pageLeft
					nums[i+1] = orderingRules[j].pageRight
				}
			}
		}
		if isSorted {
			return nums
		}
	}
}
