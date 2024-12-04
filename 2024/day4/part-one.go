package day4

func PosXMAS(pos vector, matrix [][]string) int {
    vectors := []vector{
        {x:0, y:1},
        {x:0, y:-1},
        {x:1, y:0},
        {x:1, y:1},
        {x:1, y:-1},
        {x:-1, y:0},
        {x:-1, y:1},
        {x:-1, y:-1},
    }
    aggregator := 0

    for _,vector := range vectors {
        aggregator += checkDirection("XMAS", pos, vector, matrix)
    }
    return aggregator
}
