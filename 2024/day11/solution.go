package day11

import (
	"fmt"
	"math"
	"os"
	"strings"

	"advent-of-code-2024.friess.studio/common"
)

func readInput(path string) []int {
    file, err := os.ReadFile(path)
    common.Check(err)
    
    values := strings.Fields(string(file))

    numbers := []int{}

    for _, value := range values {
        numbers = append(numbers, common.ConvertChecked(value))
    }

    return numbers
}

func accumulateStonesAfterBlinksRecursive(stones []int, blinks int) int {
    var accumulator int
    for _, stone := range stones {
        accumulator += blink(stone, blinks)
    }
    return accumulator
}

func blink(stone int, n int) int {
    if n == 0 {
        return 1
    }
    switch {
    case stone == 0:
        return blink(1, n-1)
    case int(math.Floor(math.Log10(float64(stone)) + 1)) % 2 == 0:
        base := 10
        divisor := base
        for stone / divisor > divisor {
            divisor *= base
        }
        return blink(stone / divisor, n-1) + blink(stone % divisor, n-1)
    default:
        return blink(stone * 2024, n-1)
    }
}

func accumulateStonesAfterBlinks(stones []int, blinks int) int {
	uniqueStones := map[int]int{}
	alreadySplitted := map[int][]int{}

	for _, stone := range stones {
        uniqueStones[stone] = 1
	}


	for i := 0; i < blinks; i++ {
		newUniqueStones := map[int]int{}

		for value, count := range uniqueStones {
			switch {
			case value == 0:
				newUniqueStones[1] += count
			case int(math.Floor(math.Log10(float64(value)) + 1)) % 2 == 0:
				if _, exists := alreadySplitted[value]; !exists {
					base := 10
					divisor := base
					for value/divisor > divisor {
						divisor *= base
					}
					alreadySplitted[value] = []int{value / divisor, value % divisor}
				}
				split := alreadySplitted[value]
				if split[0] > 0 {
					newUniqueStones[split[0]] += count
				}
				newUniqueStones[split[1]] += count
			default:
				newUniqueStones[value * 2024] += count
			}
		}
		uniqueStones = newUniqueStones
	}

	accumulator := 0
	for _, count := range uniqueStones {
		accumulator += count
	}

	return accumulator
}

func Solve() {
    stones := readInput("./day11/input.txt")
    fmt.Println("Part 1: ", accumulateStonesAfterBlinks(stones, 25))
    fmt.Println("Part 1: ", accumulateStonesAfterBlinks(stones, 75))
}
