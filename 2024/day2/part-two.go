package day2

func IsSafeWithProblemDampener(report []int) bool {
	if IsSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
        modifiedReport := make([]int, len(report)-1)
        copy(modifiedReport, report[:i])
        copy(modifiedReport[i:], report[i+1:])


		if IsSafe(modifiedReport) {
			return true
		}
	}

	return false
}
