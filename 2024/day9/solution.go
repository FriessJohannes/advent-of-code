package day9

import (
    "os"
    "fmt"
    "advent-of-code-2024.friess.studio/common"
)

func readInput(path string) []int {
    file, err := os.ReadFile(path)
    common.Check(err)

    file = file[:len(file)-1]

    digits := []int{}
    for _, value := range file {
        digits = append(digits, common.ConvertChecked(string(value)))
    }

    return digits
}

func Solve() {
    diskMap := readInput("./day9/input.txt")
    diskMap2 := make([]int, len(diskMap))
    copy(diskMap2, diskMap)
    fmt.Println("Part 1: ", Compress(diskMap))

    partitions := ToPartitions(diskMap2)
    ReorderFiles(partitions)
    fmt.Println("Part 2: ", CalculateChecksum(partitions))
}
