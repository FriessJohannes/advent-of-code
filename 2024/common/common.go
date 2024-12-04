package common

import "strconv"

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func ConvertChecked(str string) int {
    val, err := strconv.Atoi(str)
    Check(err)
    return val
}
