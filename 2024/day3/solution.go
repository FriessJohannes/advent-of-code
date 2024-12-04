package day3

import (
    "os"
    "fmt"

    "advent-of-code-2024.friess.studio/common"
)

type checkSafety func([]int) bool

func readInput(filename string) string {
    file, err := os.ReadFile("./day3/" + filename)
    common.Check(err)

    return string(file)
}

func scanInput(memory string, instructions bool) int {
    var (
        index = 3
        aggregator = 0
        instructionState = true
    )

    for index < len(memory) {
        if memory[index] == '(' {
            if instructions {
                instructionState = SetInstructions(&memory, &index, instructionState)
            }
            if instructionState {
                aggregator += ScanMultiplication(&memory, &index)
            }
        }
        index++;

    }

    return aggregator
}


func Solve() {
    memory := readInput("input.txt")

    fmt.Println("Part 1: ", scanInput(memory, false))
    fmt.Println("Part 2: ", scanInput(memory, true))
}

