package day1

import (
    "math"
    "advent-of-code-2024.friess.studio/common"
)



func reduceDistance(s1 []int, s2[]int) int {
    var aggregator = 0

    for i := 0; i < len(s1); i++ {
        aggregator += int(math.Abs(float64(s1[i]-s2[i])))
    }

    return aggregator
}

func SolveOne(s1 []int, s2 []int) int {
    s1 = common.MergeSort(s1)
    s2 = common.MergeSort(s2)

    return reduceDistance(s1,s2) 
}
