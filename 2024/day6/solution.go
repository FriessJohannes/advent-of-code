package day6

import (
    "fmt"
	"advent-of-code-2024.friess.studio/common"
)

type guard struct {
    position vector
    direction vector
    visited map[vector]bool
    count int
    matrix [][]string
}

type simulation struct {
    guard guard
    circles int
}

func Solve() {
    matrix := common.ReadFileToMatrix("./day6/input.txt")

    var g1 = guard{
        vector{},
        vector{x:0, y:-1},
        map[vector]bool{},
        0,
        matrix,
    }
    g1.initPosition()

    var g2 = guard{
            vector{},
            vector{x:0, y:-1},
            map[vector]bool{},
            0,
            matrix,
        }
    g2.initPosition()

    var simulation = simulation{
        g2,
        0,
    }

    fmt.Println("Part 1: ", g1.moveUntilOffMap())
    fmt.Println("Part 2: ", simulation.getCircles())
}
