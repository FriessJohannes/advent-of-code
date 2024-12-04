package day3

func convertNumber(memory *string, index *int) int {
    var num int

    *index++

    for *index < len(*memory) && (*memory)[*index] >= '0' && (*memory)[*index] <= '9' {
        num = num * 10 + int((*memory)[*index]-'0')
        *index++;
    }

    return num
}

func ScanMultiplication(memory *string, index *int) int {
    if (*memory)[*index-3:*index] == "mul" {
        num1 := convertNumber(memory, index)

        if (*memory)[*index] != ',' {
            *index++;
            return 0
        }
        num2 := convertNumber(memory, index)
        if (*memory)[*index] != ')' {
            *index++;
            return 0
        }
        return num1 * num2
    }
    return 0
}


