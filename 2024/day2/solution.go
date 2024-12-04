package day2

import (
    "os"
    "strings"
    "fmt"

    "advent-of-code-2024.friess.studio/common"
)

type checkSafety func([]int) bool

func readInput(filename string) [][]int  {
    file, err := os.ReadFile("./day2/" + filename)
    common.Check(err)

    lines := strings.Split(string(file), "\n")
    lines = lines[:len(lines)-1] // remove last \n -> empty string

    var reports [][]int

    for _, line := range lines {
        strReport := strings.Fields(line)
        var report []int
        for _, level := range strReport {
            report = append(report, common.ConvertChecked(level))
        }
        reports = append(reports, report)
    }

    return reports
}

func Solve() {
    reports := readInput("input.txt")
    fmt.Println("Part 1: ", AggregateSafeReports(reports, IsSafe))
    fmt.Println("Part 2: ", AggregateSafeReports(reports, IsSafeWithProblemDampener))
}

