package day8

import (
    "math"
    "advent-of-code-2024.friess.studio/common"
)

func captureAntinode(point common.FloatPoint, captured *map[common.IntPoint]bool) int {
        antinode := point.ToIntPoint()
        if !(*captured)[antinode] {
            (*captured)[antinode] = true
            return 1
        }
        return 0
}

func findAntinodesFor(ant1 common.IntPoint, ant2 common.IntPoint, mapRange pointRange, captured *map[common.IntPoint]bool) int {
    var (
        accumulator int = 0
        ant1f = ant1.ToFloatPoint()
        ant2f = ant2.ToFloatPoint()
    )

    vec := common.Vector{
        Start: ant1f, 
        End: ant2f,
    }

    vec.MoveTo(ant2f)
    if vec.End.IsInRange(mapRange.start, mapRange.end) {
        accumulator += captureAntinode(vec.End, captured)
    }

    vec.MoveTo(ant1f)
    vec.Rotate(math.Pi)

    if vec.End.IsInRange(mapRange.start, mapRange.end) {
        accumulator += captureAntinode(vec.End, captured)
    }
    return accumulator
}
