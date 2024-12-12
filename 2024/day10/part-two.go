package day10

func findPaths(field Field, matrix [][]int, paths *int) {
    
    directions := []Field{
        {y: -1, x: 0},
        {y: 0, x: -1},
        {y: 1, x: 0},
        {y:0, x:1},
    }

    if matrix[field.y][field.x] == 9 {
        *paths++
        return
    }

    nextValue := matrix[field.y][field.x] + 1


    for _, dir := range directions {
        y := field.y + dir.y 
        x := field.x + dir.x 
        if  y < 0 || y >= len(matrix) || x < 0 || x >= len(matrix[0]) {
            continue
        }
        if matrix[y][x] == nextValue {
            nextField := Field{y: y, x: x}
            findPaths(nextField, matrix, paths)
        }
    }
    return
}

func GetPathScore(field Field, matrix [][]int) int {
    paths := 0
    findPaths(field, matrix, &paths)
    return paths
}
