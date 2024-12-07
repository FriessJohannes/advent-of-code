package common

import (
    "os"
    "strings"
)

func ReadFileToMatrix(path string) [][]string {
    file, err := os.ReadFile(path)
    Check(err)
    lines := strings.Split(string(file), "\n")
    lines = lines[:len(lines)-1]
    var matrix [][]string 
    for _, line := range lines {
        bytesOfLine := strings.Split(line, "")
        matrix = append(matrix, bytesOfLine)
    }
    return matrix
}
