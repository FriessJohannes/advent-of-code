package day6


import (
    "math"
)

type vector struct {
    x int
    y int
}

func (guard *guard) initPosition() {
    for row := 0; row < len(guard.matrix); row++ {
        for col := 0; col < len(guard.matrix[0]); col++ {
            if guard.matrix[row][col] == "^" {
                guard.position = vector{x:col, y:row}
            }
        }
    }
}

func (vec *vector) rotate() {
    origX := vec.x
    origY := vec.y
    angle := -math.Pi/2
    vec.x = -int(float64(origX)*math.Cos(angle) - float64(origY)*math.Sin(angle))
    vec.y = -int(float64(origX)*math.Sin(angle) + float64(origY)*math.Cos(angle))
}

func (guard *guard) markMatrix(marker string) {
    if guard.matrix[guard.position.y][guard.position.x] != marker {
        guard.matrix[guard.position.y][guard.position.x] = marker
        guard.count++
    }
}

func (guard *guard) checkObstacle() bool {
    nextPos := guard.getNext()
    return guard.matrix[nextPos.y][nextPos.x] == "#"
}

func (guard *guard) isOffMap(nextPos vector) bool {
    return nextPos.x < 0 || nextPos.x > len(guard.matrix[0])-1 || nextPos.y < 0 || nextPos.y > len(guard.matrix)-1
}

func (guard *guard) getNext() vector {
    return vector{
        x:guard.position.x + guard.direction.x, 
        y:guard.position.y + guard.direction.y,
    }
}

func (guard *guard) addVisited() {
    if !guard.visited[guard.position] {
        guard.visited[guard.position] = true
        guard.count++
    }
}

func (guard *guard) moveStep() {
    for guard.checkObstacle() {
        guard.direction.rotate()
    }
    guard.position = guard.getNext()
    guard.addVisited()
    }

func (guard *guard) moveUntilOffMap() int {
    guard.addVisited()
    for {
        if guard.isOffMap(guard.getNext()) {
            return guard.count
        }
        guard.moveStep()
    }
}
