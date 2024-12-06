package day5

func checkUpdate(pages []int, rules map[int][]int) bool {
    illegals := []int{}
    for i := len(pages)-1; i >= 0; i-- {
        for _, illegal := range illegals {
            if pages[i] == illegal {
                return false
            }
        }
        illegals = append(illegals, rules[pages[i]]...)
    }
    return true
}


func AggregateMiddlePageNum(updates [][]int, rules map[int][]int) int {
    var aggregator int = 0

    for _, update := range updates {
        if checkUpdate(update, rules) {
            aggregator += update[len(update)/2]
        }
    }
    return aggregator
}
