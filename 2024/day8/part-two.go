package day8

import (
    "math"
    "advent-of-code-2024.friess.studio/common"
)

type linFunc struct{
    x float64
    k float64
    d float64
}

func findAllAntinodesFor(ant1 common.IntPoint, ant2 common.IntPoint, mapRange pointRange, captured *map[common.IntPoint]bool) int {
    var (
        accumulator int = 0
        ant1f = ant1.ToFloatPoint()
        ant2f = ant2.ToFloatPoint()
    )

    vec := common.Vector{
        Start: ant1f, 
        End: ant2f,
    }

    for vec.End.IsInRange(mapRange.start, mapRange.end) {
        accumulator += captureAntinode(vec.End, captured)
        vec.MoveTo(vec.End)
    }

    vec.MoveTo(ant1f)
    vec.Rotate(math.Pi)

    for vec.End.IsInRange(mapRange.start, mapRange.end) {
        accumulator += captureAntinode(vec.End, captured)
        vec.MoveTo(vec.End)
    }
    return accumulator
}
