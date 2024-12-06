package day5

import (
	"fmt"
	"os"
	"strings"

	"advent-of-code-2024.friess.studio/common"
)

type vector struct {
    x int
    y int
}

func splitLines(block string) []string {
    lines := strings.Split(block, "\n")
    lines = lines[:len(lines)-1]
    return lines
}


func createRules(lines []string) map[int][]int{
    orderMap := map[int][]int{}
    for _, line := range lines {
        pair := strings.Split(line, "|")
        key := common.ConvertChecked(pair[0])
        value := common.ConvertChecked(pair[1])
        orderMap[key] = append(orderMap[key], value)
    }
    return orderMap
}

func createUdateList(lines []string) [][]int {
    updates := [][]int{}

    for _, line := range lines {
        pages := []int{}
        for _, page := range strings.Split(line, ",") {
            pages = append(pages, common.ConvertChecked(page))
        }
        updates = append(updates, pages)
    }
    return updates
}

func readInput(filename string) []string {
    file, err := os.ReadFile("./day5/" + filename)
    common.Check(err)
    sections := strings.Split(string(file), "\n\n")
    return sections
}

func Solve() {
    sections := readInput("input.txt")
    rules := createRules(splitLines(sections[0]))
    updates := createUdateList(splitLines(sections[1]))

    fmt.Println("Part 1: ", AggregateMiddlePageNum(updates, rules))
    fmt.Println("Part 2: ", AggregateCorrectedMiddlePageNum(updates, rules))
}

