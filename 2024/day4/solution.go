package day4

import (
    "fmt"

    "advent-of-code-2024.friess.studio/common"
)

type vector struct {
    x int
    y int
}

type countPos func(vector, [][]string) int

func checkDirection(word string, pos vector, dir vector, matrix [][]string) int {
    var dirWord string = string(matrix[pos.y][pos.x])
    for i := 1; i < len(word); i++ {
        pos.x += dir.x
        pos.y += dir.y
        if pos.x < 0 || pos.x >= len(matrix[0]) || pos.y < 0 || pos.y >= len(matrix) {
            return 0
        }
        dirWord += matrix[pos.y][pos.x]
        if dirWord != word[:len(dirWord)] {
            return 0
        }
    }
    return 1
}

func findWord(start string, matrix [][]string, count countPos) int {
    height := len(matrix)
    width := len(matrix[0])
    aggregator := 0

    for row := 0; row < height; row++ {
        for col := 0; col < width; col++ {
            if matrix[row][col] == start {
                pos := vector{
                    x: col,
                    y: row,
                }
                aggregator += count(pos, matrix)
            }
        }
    }
    return aggregator
}

func Solve() {
    matrix := common.ReadFileToMatrix("./day4/input.txt")

    fmt.Println("Part 1: ", findWord("X", matrix, PosXMAS))
    fmt.Println("Part 2: ", findWord("A", matrix, PosX_MAS))
}

