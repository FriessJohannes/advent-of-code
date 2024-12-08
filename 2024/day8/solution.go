package day8

import (
	"fmt"

	"advent-of-code-2024.friess.studio/common"
)

type evalFunc func(point1 common.IntPoint, point2 common.IntPoint, mapRange pointRange, captured *map[common.IntPoint]bool) int

type pointRange struct {
    start common.FloatPoint
    end common.FloatPoint
}

func isAlphaNumeric(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func getAntennaCoordinates(matrix [][]string) map[string][]common.IntPoint {
    antennas := map[string][]common.IntPoint{}

    for row := 0; row < len(matrix); row++ {
        for col := 0; col < len(matrix[0]); col++ {
            freq := matrix[row][col]
            if isAlphaNumeric(freq[0]) { 
                antennas[freq] = append(antennas[freq], common.IntPoint{X: col, Y: row})
            }
        }
    }
    return antennas
}

func accumulateAntinodes(freq map[string][]common.IntPoint, mapRange pointRange, fn evalFunc) int {
    var (
        captured = map[common.IntPoint]bool{}
        accumulator int = 0
    )

    for _, antennas := range freq {
        for _, ant1 := range antennas {
            for _, ant2 := range antennas {
                if  ant1.IsEqual(ant2) {
                    continue
                }
                accumulator += fn(ant1, ant2, mapRange, &captured)
            }
        }
    }
    return accumulator
}



func Solve() {
    matrix := common.ReadFileToMatrix("./day8/input.txt")
    frequencies := getAntennaCoordinates(matrix)
    mapRange := pointRange{
        start: common.FloatPoint{X:0, Y:0},
        end: common.FloatPoint{X:float64(len(matrix[0])), Y:float64(len(matrix))},
    }
    fmt.Println("Part 1: ", accumulateAntinodes(frequencies, mapRange, findAntinodesFor))
    fmt.Println("Part 2: ", accumulateAntinodes(frequencies, mapRange, findAllAntinodesFor))
}
