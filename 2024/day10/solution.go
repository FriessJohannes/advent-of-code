package day10

import (
    "fmt"
    "advent-of-code-2024.friess.studio/common"
)

type Field struct {
    y int
    x int
}

type getScore func(field Field, matrix [][]int) int

func accumulateScores(matrix [][]int, getScore getScore) int {
    var accumulator int = 0
    for row := 0; row < len(matrix); row++ {
        for col := 0; col < len(matrix[0]); col++ {
            if matrix[row][col] == 0 {
                accumulator += getScore(Field{y: row, x:col}, matrix)
            }
        }
    }
    return accumulator
}

func Solve() {
    matrix := common.ReadFileToIntMatrix("./day10/input.txt")
    fmt.Println("Part 1: ", accumulateScores(matrix, GetTopsScore))
    fmt.Println("Part 2: ", accumulateScores(matrix, GetPathScore))
}
