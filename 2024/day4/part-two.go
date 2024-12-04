package day4

type diagonal struct {
    direction vector
    start vector
}

func PosX_MAS(pos vector, matrix [][]string) int {
    if pos.x == 0 || pos.y == 0 || pos.x == len(matrix[0])-1 || pos.y == len(matrix)-1 {
        return 0
    }
    diagonals := []diagonal{
        {direction:vector{x:1, y:1}, start:vector{x:pos.x-1, y:pos.y-1}},
        {direction:vector{x:-1, y:1}, start:vector{x:pos.x+1, y:pos.y-1}},
    }
    
    aggregator := 0

    for _, diagonal := range diagonals {
        aggregator += checkDirection("MAS", diagonal.start, diagonal.direction, matrix)
        aggregator += checkDirection("SAM", diagonal.start, diagonal.direction, matrix)
    }
    
    if aggregator == 2 {
        return 1
    } else {
        return 0
    }
}

