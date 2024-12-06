package day5

type Correction struct {
        pages []int
        before map[int][]int
        after map[int][]int
        valueBefore map[int][]int
    }

func topologicalSort(update []int, rules map[int][]int) []int {
    pageSet := makeSet(update)
    graph := buildGraph(pageSet, rules)
    inDegree := calculateInDegrees(pageSet, graph)
    return performTopologicalSort(pageSet, graph, inDegree)
}

func makeSet(pages []int) map[int]bool {
    pageSet := make(map[int]bool, len(pages))
    for _, p := range pages {
        pageSet[p] = true
    }
    return pageSet
}

func buildGraph(pageSet map[int]bool, rules map[int][]int) map[int][]int {
    graph := make(map[int][]int, len(pageSet))
    for p := range pageSet {
        graph[p] = []int{}
    }

    for p := range pageSet {
        successors := rules[p]
        for _, s := range successors {
            if pageSet[s] {
                graph[p] = append(graph[p], s)
            }
        }
    }
    return graph
}

func calculateInDegrees(pageSet map[int]bool, graph map[int][]int) map[int]int {
    inDegree := make(map[int]int, len(pageSet))
    for p := range pageSet {
        inDegree[p] = 0
    }

    for _, successors := range graph {
        for _, s := range successors {
            inDegree[s]++
        }
    }

    return inDegree
}

func performTopologicalSort(pageSet map[int]bool, graph map[int][]int, inDegree map[int]int) []int {
    queue := []int{}
    for p := range pageSet {
        if inDegree[p] == 0 {
            queue = append(queue, p)
        }
    }

    var order []int
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        order = append(order, current)

        for _, neigh := range graph[current] {
            inDegree[neigh]--
            if inDegree[neigh] == 0 {
                queue = append(queue, neigh)
            }
        }
    }
    if len(order) != len(pageSet) {
        return []int{}
    }

    return order
}

func AggregateCorrectedMiddlePageNum(updates [][]int, beforeToAfter map[int][]int) int {
    var aggregator int = 0

    for _, update := range updates {
        if !checkUpdate(update, beforeToAfter) {
            correctedUpdate := topologicalSort(update, beforeToAfter)
            aggregator += correctedUpdate[len(update)/2]
        }
    }
    return aggregator
}
