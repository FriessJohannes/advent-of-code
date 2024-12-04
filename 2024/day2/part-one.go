package day2

func IsSafe(report []int) bool {
	increasing := false
	decreasing := false

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff > 3 || diff < -3 || diff == 0 {
			return false
		}
		if diff > 0 {
			increasing = true
		}
		if diff < 0 {
			decreasing = true
		}
	}
	return !(increasing && decreasing)
}

func AggregateSafeReports(reports [][]int, check checkSafety) int {
    var aggregator int

    for _, report := range reports {
        if check(report) {
            aggregator++;
        }
    }

    return aggregator
}
