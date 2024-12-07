package day6

// 1654 < solution < 1696

//import "fmt"

type guardState struct {
    position vector
    direction vector
} 

func (guard *guard) isInCircle() bool {
    visited := map[guardState]bool{}

    startState := guardState{guard.position, guard.direction}
    visited[startState] = true
    
    for {
        guard.moveStep()
        
        currentState := guardState{guard.position, guard.direction}

        if visited[currentState] {
            return true
        }

        visited[currentState] = true

        if guard.isOffMap(guard.getNext()) {
            return false
        }
    }
}

func (guard *guard) restoreState(state guardState) {
    guard.position = state.position
    guard.direction = state.direction
}

func (sim *simulation) start() {
    matrix := sim.guard.matrix
    obstaclePos := sim.guard.getNext()
    new := guard{
        sim.guard.position,
        sim.guard.direction,
        map[vector]bool{},
        0,
        matrix,
    }
    matrix[obstaclePos.y][obstaclePos.x] = "#"
    if new.isInCircle() {
        sim.circles++
    }
    matrix[obstaclePos.y][obstaclePos.x] = "."
}

func (sim *simulation) getCircles() int {
    sim.guard.addVisited()
    for {
        if sim.guard.isOffMap(sim.guard.getNext()) {
            return sim.circles
        }
        for sim.guard.checkObstacle() {
            sim.guard.direction.rotate()
        }
        if !sim.guard.visited[sim.guard.getNext()] {
            sim.start()
        }
        sim.guard.moveStep()
    }
}
