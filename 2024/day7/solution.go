package day7

import (
    "os"
    "strings"
    "fmt"
	"advent-of-code-2024.friess.studio/common"
)

type Operation func(prevVal, currentVal int) int

type equation struct {
    result int
    values []int
}

func readInput(path string) []equation {
    file,err := os.ReadFile(path)
    common.Check(err)

    lines := strings.Split(string(file),"\n")
    lines = lines[:len(lines)-1]

    equations := []equation{}

    for _, line := range lines {
        splitLine := strings.Split(line, ":")
        strValues := strings.Fields(splitLine[1])
        intValues := []int{}
        for _, value := range strValues {
            intValues = append(intValues, common.ConvertChecked(value))
        }
        equations = append(equations, equation{
            result: common.ConvertChecked(splitLine[0]),
            values: intValues,
        })
    }

    return equations
}

func AddOperation(prevVal, currentVal int) int {
    return prevVal + currentVal
}

func MultiplyOperation(prevVal, currentVal int) int {
    return prevVal * currentVal
}

func ConcatOperation(prevVal, currentVal int) int {
    var padding int = 1
    for padding <= currentVal {
        padding *= 10
    }
    return prevVal*padding + currentVal
}

func isValid(eq equation, operations []Operation) bool {
    n := len(eq.values)
    if n == 0 {
        return eq.result == 0
    }

    dp := make([]map[int]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = map[int]bool{}
    }

    dp[0][eq.values[0]] = true

    for i := 1; i < n; i++ {
        for prevVal := range dp[i-1] {
            for _, op := range operations {
                newVal := op(prevVal, eq.values[i])
                if newVal <= eq.result {
                    dp[i][newVal] = true
                }
            }
        }
    }

    return dp[n-1][eq.result]
}

func accumulateValidEq(equations []equation, operations []Operation) int {
    var accumulator int
    for _, eq := range equations {
        if isValid(eq, operations) {
            accumulator += eq.result
        }
    }
    return accumulator
}

func Solve() {

    equations := readInput("./day7/input.txt")

    fmt.Println("Part 1: ", accumulateValidEq(equations, []Operation{AddOperation, MultiplyOperation}))
    fmt.Println("Part 2: ", accumulateValidEq(equations, []Operation{AddOperation, MultiplyOperation, ConcatOperation}))
}
