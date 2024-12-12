package day10

func findTops(field Field, matrix [][]int, tops map[Field]bool) {
    
    directions := []Field{
        {y: -1, x: 0},
        {y: 0, x: -1},
        {y: 1, x: 0},
        {y:0, x:1},
    }

    if matrix[field.y][field.x] == 9 {
        tops[field] = true
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
            findTops(nextField, matrix, tops)
        }
    }
    return
}

func GetTopsScore(field Field, matrix [][]int) int {
    tops := map[Field]bool{}
    findTops(field, matrix, tops)
    return len(tops)
}
