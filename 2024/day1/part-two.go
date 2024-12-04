package day1

import (
	"advent-of-code-2024.friess.studio/common"
)

func calcSimilarityScore(s1 []int, s2 []int) int {
    var (
        aggregator = 0
        ps2 = 0
        current = s1[0]
        count = 0
    )

    for _, value := range s1 {
        if value != current {
            count = 0
            current = value
            for ps2 < len(s2) && s2[ps2] <= value {
                if s2[ps2] == value {
                    count++;
                }
                ps2++;
            }
        }
        aggregator += value * count
    }

    return aggregator
}

func SolveTwo(s1 []int, s2 []int) int {
    s1 = common.MergeSort(s1)
    s2 = common.MergeSort(s2)
    return calcSimilarityScore(s1, s2)
}
