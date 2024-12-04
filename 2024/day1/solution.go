package day1

import (
	"fmt"
	"os"
	"strings"
	"unicode"

    "advent-of-code-2024.friess.studio/common"
)

func readInput(filename string) ([]int, []int) {
    file, err := os.ReadFile("./day1/" + filename)
    common.Check(err)

    f := func(c rune) bool {
		return unicode.IsSpace(c)
	}

    fields := strings.FieldsFunc(string(file), f)

    var (
        s1 = []int{}
        s2 = []int{}
    )

    for index, field := range fields {
        if (index % 2 == 0) {
            s1 = append(s1, common.ConvertChecked(field))
        } else {
            s2 = append(s2, common.ConvertChecked(field))
        }
    }
    return s1, s2
}

func Solve() {
    s1, s2 := readInput("input.txt")
    fmt.Println("Part 1: ", SolveOne(s1, s2))
    fmt.Println("Part 2: ", SolveTwo(s1, s2))
}
